package service

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type CreatingFilesBackEnd struct {
	path         string
	completePath string
}

func (c *CreatingFilesBackEnd) creatingFilesBackEnd() {

	dir := fmt.Sprintf("%s/backend", c.path)
	err := os.MkdirAll(dir, os.ModePerm)

	if err != nil {
		log.Fatal(err)
		return
	}

}

func (c *CreatingFilesBackEnd) creatingStructureBase() {
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
		err := os.MkdirAll(fmt.Sprintf("%s/%s", c.completePath, dir), os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("estrutura criada com sucesso!")
	}

}

func (c *CreatingFilesBackEnd) AirToml() string {
	return `root = "."
testdata_dir = "testdata"
tmp_dir = "tmp"

[build]
  args_bin = []
  bin = "./tmp/main"
  cmd = "go build -o ./tmp/main ."
  delay = 1000
  exclude_dir = ["assets", "tmp", "vendor", "testdata"]
  exclude_file = []
  exclude_regex = ["_test.go"]
  exclude_unchanged = false
  follow_symlink = false
  full_bin = ""
  include_dir = []
  include_ext = ["go", "tpl", "tmpl", "html"]
  include_file = []
  kill_delay = "0s"
  log = "build-errors.log"
  poll = false
  poll_interval = 0
  post_cmd = []
  pre_cmd = []
  rerun = false
  rerun_delay = 500
  send_interrupt = false
  stop_on_error = false

[color]
  app = ""
  build = "yellow"
  main = "magenta"
  runner = "green"
  watcher = "cyan"

[log]
  main_only = false
  time = false

[misc]
  clean_on_exit = false

[proxy]
  app_port = 0
  enabled = false
  proxy_port = 0

[screen]
  clear_on_rebuild = false
  keep_scroll = true
`
}

func (c *CreatingFilesBackEnd) TailwindReactJS() string {
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

	return data

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

func (c *CreatingFilesBackEnd) InjectConfigurationTailwind(framework string) error {
	path := fmt.Sprintf("%s/%s", c.path, "tailwind.config.js")
	f, err := os.Open(path)
	if err != nil {
		return err
	}

	defer f.Close()

	w := bufio.NewWriter(f)

	bf, err := w.WriteString(framework)

	defer w.Flush()

	if err != nil {
		return err
	}

	fmt.Printf("bufferzing tailwind.config.js %d bytes", bf)

	return nil

}
