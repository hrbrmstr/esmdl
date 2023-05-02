package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/urfave/cli/v2"
)

func main() {

	app := &cli.App{
		Name:  "download-esm",
		Usage: "Download ESM modules from npm and jsdelivr",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "package",
				Aliases:  []string{"p"},
				Usage:    "Package to download",
				Required: true,
			},
			&cli.StringFlag{
				Name:        "location",
				Aliases:     []string{"l"},
				Usage:       "Location to save files",
				DefaultText: ".",
			},
		},

		Action: func(c *cli.Context) error {

			packageArg := c.String("package")
			location := c.String("location")

			root := filepath.Join(location)
			if _, err := os.Stat(root); os.IsNotExist(err) {
				os.MkdirAll(root, os.ModePerm)
			}

			var code string

			if strings.HasPrefix(packageArg, "https://") {

				response, err := http.Get(packageArg)
				if err != nil {
					return err
				}

				defer response.Body.Close()

				bytes, err := io.ReadAll(response.Body)
				if err != nil {
					return err
				}

				code = string(bytes)

			} else {

				code = fetchCode(packageArg)

			}

			originalFile := extractOriginalFile(code)
			path := simplifyPath(originalFile)

			// Rewrite code and save to path
			rewrittenCode, capturedPaths := rewriteCode(code)
			err := os.WriteFile(filepath.Join(root, path), []byte(rewrittenCode), 0644)
			if err != nil {
				return err
			}

			fmt.Fprintln(os.Stderr, path)

			// Do the same thing for all the captured paths, recursively
			toFetch := map[string]string{}
			for k, v := range capturedPaths {
				toFetch[k] = v
			}

			for len(toFetch) > 0 {

				for path, simplifiedPath := range toFetch {

					delete(toFetch, path)

					url := "https://cdn.jsdelivr.net" + path
					response, err := http.Get(url)
					if err != nil {
						return err
					}

					defer response.Body.Close()
					bytes, err := io.ReadAll(response.Body)
					if err != nil {
						return err
					}
					code := string(bytes)

					rewrittenCode, moreCapturedPaths := rewriteCode(code)
					err = os.WriteFile(filepath.Join(root, simplifiedPath), []byte(rewrittenCode), 0644)
					if err != nil {
						return err
					}

					fmt.Fprintln(os.Stderr, simplifiedPath)

					for k, v := range moreCapturedPaths {
						toFetch[k] = v
					}

				}
			}

			return nil

		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

}

func fetchCode(packageArg string) string {

	url := fmt.Sprintf("https://cdn.jsdelivr.net/npm/%s/+esm", packageArg)
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	bytes, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	return string(bytes)

}

func extractOriginalFile(content string) string {

	pattern := `Original file: (\/npm\/[^\s]+)`
	re := regexp.MustCompile(pattern)
	match := re.FindStringSubmatch(content)

	if len(match) > 1 {
		return match[1]
	} else {
		panic("Could not find original file")
	}

}

func simplifyPath(path string) string {

	split := strings.Split(path, "/npm/")
	split[1] = strings.TrimPrefix(split[1], "@")

	packageInfo := strings.SplitN(split[1], "@", 2)

	packageName := strings.ReplaceAll(packageInfo[0], "/", "-")

	packageVersion := strings.Split(packageInfo[1], "/")[0]
	packageVersion = strings.ReplaceAll(packageVersion, ".", "-")

	simplifiedName := fmt.Sprintf("%s-%s.js", packageName, packageVersion)

	return simplifiedName
}

func rewriteCode(code string) (string, map[string]string) {

	pattern := `(?P<keyword>import|export)\s*(?P<imports>\{?[^}]+?\}?)\s*from\s*"(?P<path>\/npm\/[^"]+)"`
	re := regexp.MustCompile(pattern)
	capturedPaths := map[string]string{}

	replaceImport := func(match string) string {

		submatch := re.FindStringSubmatch(match)

		keyword := submatch[1]
		imports := submatch[2]
		path := submatch[3]

		simplifiedPath := simplifyPath(path)

		capturedPaths[path] = simplifiedPath

		return fmt.Sprintf(`%s %s from "./%s";`, keyword, imports, simplifiedPath)
	}

	rewrittenCode := re.ReplaceAllStringFunc(code, replaceImport)
	rewrittenCode = removeSourceMappingComments(rewrittenCode)

	return rewrittenCode, capturedPaths

}

func removeSourceMappingComments(code string) string {

	pattern := `\/\/#\s*sourceMappingURL=.*?\.map`
	re := regexp.MustCompile(pattern)
	cleanedCode := re.ReplaceAllString(code, "")

	return cleanedCode

}
