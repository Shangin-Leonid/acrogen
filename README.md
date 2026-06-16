# :bulb: acrogen

![Go version](https://img.shields.io/github/go-mod/go-version/Shangin-Leonid/acrogen?style=flat&logo=Go&logoColor=E7FEFB&logoSize=auto&labelColor=00ADD8&color=00ADD8)
![Makefile](https://img.shields.io/static/v1?label=&message=Makefile&style=flat&logo=CMake&color=FE7A16)
![Linux](https://img.shields.io/static/v1?label=&message=Linux&style=flat&logo=linux&color=2B822F)
![Windows](https://img.shields.io/static/v1?label=&message=Windows&style=flat&logo=windows&color=4169e1)

**acrogen** (Acronym Generator) is a fast and efficient Command Line Interface (CLI) tool written in Go, designed to generate acronyms and abbreviations from user query.

---

## ✨ Features

- :ru: **Russian language is supported (for generation, not for UI)**
- :uk: **English language is supported (UI, generation)**
- :desktop_computer: **Intuitive CLI**
- 🛠️ **Easy to build and run with `makefile`**
- ⚡ **High performance**
- 📦 **Modular Architecture:**
- :twisted_rightwards_arrows: **Concurrency is used**

---

## 🚀 Getting started

### Prerequisites
To build and run this project, you need:
- **Go** (see version above)
- **Make** (optional, for quick build and run)

### Installation and setup

1. :arrow_down: Clone the repository:
```bash
git clone https://github.com
cd acrogen
```

2. :gear: Compile the project using the `Makefile`:
```bash
make build
```
*This will generate an executable binary in the root directory.*

---

## :keyboard: Usage

Run the compiled binary by passing a string or a text source to process:

```bash
./acrogen "As Far As I Know"
# Output: AFAIK
```

*(Optional: If your tool supports CLI flags, you can document them below)*
```bash
./acrogen --help
```

---

## :open_file_folder: Project structure

The repository follows a clean and structured package layout:
- `makefile` — Make commands to build and run quickly
- `ag/` — core of acronym representation and generation
- `algo/` — common algorithms used for generation
- `cont/` — some custom containers
- `data/` — source data, russian and english dictionaries, default storage for log and output files
- `fio/` — file Input/Output subsystem (reading source data and writing results)
- `ui/` — user CLI, interaction components
- `utils/` — common helpers and utility functions

---

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

## 🤝 Feedback
**You can contact with me using info from [my Github profile](https://github.com/Shangin-Leonid/)**

**If you find 'acrogen' helpful, please give me a star :star:**
