package service

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type CreatingFilesBackEnd struct {
	path         string
	CompletePath string
}

func (c *CreatingFilesBackEnd) creatingFilesBackEnd() {

	dir := fmt.Sprintf("%s/backend", c.path)
	err := os.MkdirAll(dir, os.ModePerm)

	if err != nil {
		log.Fatal(err)
		return
	}

}

func (c *CreatingFilesBackEnd) CreatingStructureBase() {
	projectDir := []string{
		"cmd",
		"pkg",
		"internal",
		"api",
		"configs",
		"data",
		"scripts",
		"docs",
		"test",
	}

	for _, dir := range projectDir {
		err := os.MkdirAll(fmt.Sprintf("%s/%s", c.CompletePath, dir), os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("estrutura criada com sucesso!")
	}

}

func (c *CreatingFilesBackEnd) TailwindReactJS() {
	data := `/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./src/**/*.{js,jsx,ts,tsx}",
  ],
  theme: {
    extend: {},
  },
  plugins: [],
}`

	c.injectConfigurationTailwind(data)

}

func (c *CreatingFilesBackEnd) TailwindVueJS() {
	data := `module.exports = {
  content: [
    "./index.html",
    "./src/**/*.{vue,js,ts,jsx,tsx}",
  ],
  theme: {
    extend: {},
  },
  plugins: [],
}`

	c.injectConfigurationTailwind(data)

}

func (c *CreatingFilesBackEnd) CreateFileBackend(filename, file string) error {
	pathFile := fmt.Sprintf("%s/backend/%s", c.path, filename)
	_, err := os.Create(pathFile)
	if err != nil {
		return err
	}

	err = os.WriteFile(pathFile, []byte(file), 0666)

	if err != nil {
		return err
	}

	return nil
}

func (c *CreatingFilesBackEnd) injectConfigurationTailwind(data string) error {
	path := fmt.Sprintf("%s/%s", c.path, "tailwind.config.js")
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	w := bufio.NewWriter(f)
	bf, err := w.WriteString(fmt.Sprintf("%s\n", data))
	if err != nil {
		return err
	}
	fmt.Println(path)
	fmt.Printf("buffering tailwind.config.js %d bytes\n", bf)
	err = w.Flush()
	if err != nil {
		return err
	}
	return nil
}

func (c *CreatingFilesBackEnd) injectTailwindAnnotationsCss(cssFileName string) error {
	data := `@tailwind base;
@tailwind components;
@tailwind utilities;`

	path := fmt.Sprintf("%s/src/%s", c.path, cssFileName)
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}

	defer f.Close()

	w := bufio.NewWriter(f)

	bf, err := w.WriteString(fmt.Sprintf("%s\n", data))

	if err != nil {
		return err
	}

	fmt.Println(path)
	fmt.Printf("buffering %s %d bytes\n", cssFileName, bf)

	err = w.Flush()

	if err != nil {
		return err
	}

	return nil

}
