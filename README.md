# Pokédex CLI

Welcome to the Pokédex CLI, a command-line interface developed in Go for exploring the world of Pokémon. This tool allows you to search, view, and manage Pokémon data directly from your terminal, bringing the extensive Pokédex into your command line environment.
<br>
<br>


## Features

Pokédex CLI offers a rich set of features designed for Pokémon enthusiasts and trainers alike, enabling exploration, collection, and management of Pokémon directly from the command line:

- **Explore Various Locations**: Navigate through different Pokémon regions and discover Pokémon in various location areas.
- **Catch Pokémon**: Embark on exciting adventures to catch Pokémon and add them to your Pokédex.
- **Inspect Your Catches**: Get detailed information on your captured Pokémon, including stats, abilities, and more.
- **Manage Your Pokédex**: View and manage your collection of Pokémon, tracking your progress as a Pokémon Trainer.
- **Helpful Navigation**: Easily navigate through the application with helpful commands for exploring previous and next location areas.
<br>

## Getting Started

### Prerequisites

Before you begin, ensure you have Go installed on your system. The project is developed with Go, so you'll need it to run and possibly contribute to the project. If you don't have Go installed, you can download it from the official Go website.
<br>
<br>
### Installation

To install Pokédex CLI, clone the repository to your local machine using the following command:

```bash
git clone https://github.com/cosmindevelops/pokedexcli.git
```

Navigate into the project directory:

```bash
cd pokedexcli
```

Build the project using:

```bash
go build
```
<br>

## Usage

Here are some examples of how you can use the commands in Pokédex CLI to explore the world of Pokémon:

- **Getting Help**
  To view a list of all available commands and their descriptions:
  ```bash
  ./pokedexcli help
  ```

- **Exploring Location Areas**
  To list the next page of location areas where Pokémon can be found:
  ```bash
  ./pokedexcli map
  ```
  To navigate back to the previous page of location areas:
  ```bash
  ./pokedexcli mapb
  ```

- **Finding and Catching Pokémon**
  To explore a specific location and see the Pokémon available there:
  ```bash
  ./pokedexcli explore [location_area]
  ```
  To attempt to catch a Pokémon and add it to your Pokédex:
  ```bash
  ./pokedexcli catch [pokemon_name]
  ```

- **Inspecting Caught Pokémon**
  For viewing detailed information about a caught Pokémon:
  ```bash
  ./pokedexcli inspect [pokemon_name]
  ```

- **Viewing Your Pokédex**
  To see all the Pokémon you've caught and added to your Pokédex:
  ```bash
  ./pokedexcli pokedex
  ```

- **Exiting the Program**
  When you're done and wish to exit Pokédex CLI:
  ```bash
  ./pokedexcli exit
  ```
Replace `[location_area]` and `[pokemon_name]` with the actual location or Pokémon name you're interested in. This approach helps users understand not just what they can do but exactly how to do it.
<br>
<br>


