# IMDd-searcher allows to search the [IMDB-dataset](https://www.imdb.com/interfaces/) tables for information.
## Currently the following search functions are available:
- FindInfoByPersonName: allows to find information about the person and basic information about all titles in which he was seen.
- FindTitleAndCastInfoByTitleName: allows to find information about the film by its title, including its rating and the actors who took part in it.
- FindTitlesByPersonName: allows to find information about all titles by the person's name.
- FindAllTitlesBySpecificYear: allows to basic information about all titles shot in a particular year.
## Build and run the program:
**Preconditions: you must place all [IMDb database files](https://datasets.imdbws.com/) in the database folder of the current project directory.**
**Reminder: to build the application correctly, it needs to be placed in GOPATH/src/github.com/**
- Windows: To build and run a program, simply run run.bat.
- Ubuntu:  To build and run a program, simply run "make" command.