# :bulb: acrogen

![Go version](https://img.shields.io/github/go-mod/go-version/Shangin-Leonid/acrogen?style=flat&logo=Go&logoColor=E7FEFB&logoSize=auto&labelColor=00ADD8&color=00ADD8)
![Makefile](https://img.shields.io/static/v1?label=&message=Makefile&style=flat&logo=CMake&color=FE7A16)
![Linux](https://img.shields.io/static/v1?label=&message=Linux&style=flat&logo=linux&color=2B822F)
![Windows](https://img.shields.io/static/v1?label=&message=Windows&style=flat&logo=windows&color=4169e1)

**acrogen** (Acronym Generator) is a fast and efficient Command Line Interface (CLI) tool written in Go, designed to generate acronyms and abbreviations from user query.

---

## ✨ Features

- :desktop_computer: **Intuitive CLI**
- 🛠️ **Easy to build and run with `makefile`**
- ⚡ **High performance**
- 📦 **Modular Architecture:**
- :twisted_rightwards_arrows: **Concurrency supporting is going to be added**

---

## :open_file_folder: Project Structure

The repository follows a clean and structured package layout:
- `algo/` — Core algorithms and acronym generation logic.
- `fio/` — File Input/Output subsystem (reading source text and writing results).
- `ui/` — Terminal-based user interface and interaction components.
- `utils/` — Common helpers and utility functions.

---

## 🚀 Getting Started

### Prerequisites
To build and run this project, you need:
- **Go** (version 1.18 or higher)
- **Make** (optional, for using the Makefile)

### Installation and Setup

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

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

## 🤝 Feedback
**You can contact with me using info from [my Github profile](https://github.com/Shangin-Leonid/)**

**If you find 'acrogen' helpful, please give me a star :star:**

