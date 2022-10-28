package piston

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/ChampManZ/ExeCode/v2/internal/utils"
)

type Package struct {
	Language string `json:"language"`
	Version  string `json:"version"`
}

type byLanguage []Package

func (p byLanguage) Len() int           { return len(p) }
func (p byLanguage) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p byLanguage) Less(i, j int) bool { return p[i].Language < p[j].Language }

func NewPackage(language string, version string) Package {
	return Package{
		language,
		version,
	}
}

// ReadPackageFile reads a newline seperated list of piston packages
// of the form "language:version" and returns a list of Package
func ReadPackageFile(path string) ([]Package, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	pistonPackages := []Package{}
	for scanner.Scan() {
		packageInfo := strings.Split(scanner.Text(), ":")
		p := NewPackage(packageInfo[0], packageInfo[1])
		pistonPackages = append(pistonPackages, p)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return pistonPackages, nil
}

func EnsurePackagesFromFile(path string, pistonClient *Client) error {

	type Result struct {
		PistonPackages []Package
		Err            error
	}
	c := make(chan Result)
	go func() {
		pistonPackages, err := pistonClient.GetInstalledPackages()
		c <- Result{PistonPackages: pistonPackages, Err: err}
	}()

	requiredPackages, err := ReadPackageFile(path)
	if err != nil {
		return err
	}
	result := <-c
	if result.Err != nil {
		return result.Err
	}
	installedPackages := result.PistonPackages

	sort.Sort(byLanguage(requiredPackages))
	sort.Sort(byLanguage(installedPackages))
	if utils.ArrayEqual(requiredPackages, installedPackages) {
		return nil
	}

	min := 0
	for _, pp := range requiredPackages {
		i := utils.InArray(pp, installedPackages[min:])
		if i == -1 {
			// TODO: Install piston package here
			fmt.Println("Missing package: ", pp)
			fmt.Println("installing...")
			err := pistonClient.InstallPackage(pp)
			if err != nil {
				return err
			}
		} else {
			min = i
			fmt.Println("Found package: ", pp)
		}
	}
	return nil
}