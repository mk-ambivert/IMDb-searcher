IMDb Dataset Details taken from website https://www.imdb.com/interfaces/ and expanded with **Notes** for features of each table.


## Dataset details
Each dataset is contained in a gzipped tab-separated-values (TSV) formatted
file in the UTF-8 character set. The first line in each file contains headers
that describe what is in each column. A "\\N" is used to denote that a
particular field is missing or has a NULL value for that title or name. It
should be noted that the data available for download from the IMDb website is
not the full dataset, but it will suffice for our purposes. The available IMDb
data files are as follows:


### name.basics.tsv.gz
Contains the following information for names:

  - nconst (string) - alphanumeric unique identifier of the name/person.
  - primaryName (string)– name by which the person is most often credited.
  - birthYear – in YYYY format.
  - deathYear – in YYYY format if applicable, else "\\N".
  - primaryProfession (array of strings) – the top-3 professions of the person.
  - knownForTitles (array of tconsts) – titles the person is known for.
____


### title.basics.tsv.gz
Contains the following information for titles:

  - tconst (string) - alphanumeric unique identifier of the title.
  - titleType (string) – the type/format of the title (e.g. movie, short,
    tvseries, tvepisode, video, etc).
  - primaryTitle (string) – the more popular title / the title used by the
  filmmakers on promotional materials at the point of release.
  - originalTitle (string) - original title, in the original language.
  - isAdult (boolean) - 0: non-adult title; 1: adult title.
  - startYear (YYYY) – represents the release year of a title. In the case of TV
  Series, it is the series start year.
  - endYear (YYYY) – TV Series end year. "\\N" for all other title types.
  - runtimeMinutes – primary runtime of the title, in minutes.
  - genres (string array) – includes up to three genres associated with the
  title.
____


### title.akas.tsv.gz
Contains the following information for titles:

  - titleId (string) - a tconst which is an alphanumeric unique identifier of
  the title.
  - ordering (integer) – a number to uniquely identify rows for a given titleId.
  - title (string) – the localised title.
  - region (string) - the region for this version of the title.
  - language (string) - the language of the title.
  - types (array) - Enumerated set of attributes for this alternative title. One
  or more of the following: "alternative", "dvd", "festival", "tv", "video",
  "working", "original", "imdbDisplay". New values may be added in the future
  without warning.
  - attributes (array) - Additional terms to describe this alternative title,
  not enumerated.
  - isOriginalTitle (boolean) – 0: not original title; 1: original title.
  
  **Notes:** 
  - titleId (string) - In reality unique title identifiers are placed in increasing order and
    **lines can be duplicated(have the same titleId(title tconst))**
	  **because the same titleId(title tconst) can have different titles**
	  **depending on the localisation and the region.**
  - types (array)
  - attributes (array) 
    – Note that the types and attributes are referred to as an array,
    but this is not true. There seems to be only one line
	  for each titleId and order pair. Many of these fields contain "\\N".
  - example:
  ```
  titleId	ordering	title	region	language	types	attributes	isOriginalTitle
  tt13683856	1	Под землёй	RU	\N	imdbDisplay	\N	0
  tt13683856	2	Storgetnya	\N	\N	original	\N	1
  tt13683856	3	Storgetnya	AM	\N	imdbDisplay	\N	0
  tt13683856	4	Storgetnya	FR	\N	imdbDisplay	\N	0
  ```
____


### title.crew.tsv.gz
Contains the director and writer information for all the titles in IMDb. Fields
include:

  - tconst (string) - alphanumeric unique identifier of the title.
  - directors (array of nconsts) - director(s) of the given title.
  - writers (array of nconsts) – writer(s) of the given title.
  
  **Notes:**
  - directors (array of nconsts) - Many of these fields contain "\N".
  - writers (array of nconsts) – Most of these fields contain "\N".
____


### title.episode.tsv.gz
Contains the tv episode information. Fields include:

  - tconst (string) - alphanumeric identifier of episode.
  - parentTconst (string) - alphanumeric identifier of the parent TV Series.
  - seasonNumber (integer) – season number the episode belongs to.
  - episodeNumber (integer) – episode number of the tconst in the TV series.
  
  **Notes:**
  - parentTconst (string) - It may appear that all the available episode IDs(tconst) 
    of one of the "the parent TV Series"(parentTconst) are clustered together, but this is incorrect.
  - example:
  ```
  tconst	parentTconst	seasonNumber	episodeNumber
  tt9496884	tt0316967	\N	\N
  tt9496888	tt2143123	8	7
  tt9496892	tt0316967	\N	\N
  tt9496894	tt0316967	\N	\N
  tt9496896	tt0316967	\N	\N
  ```
____


### title.principals.tsv.gz
Contains the principal cast/crew for titles

  - tconst (string) - alphanumeric unique identifier of the title.
  - ordering (integer) – a number to uniquely identify rows for a given titleId.
  - nconst (string) - alphanumeric unique identifier of the name/person.
  - category (string) - the category of job that person was in.
  - job (string) - the specific job title if applicable, else "\\N".
  - characters (string) - the name of the character played if applicable, else "\\N" 
  
  **Notes:**
  - tconsts (string) - In reality unique identifiers are placed in increasing order and
    **lines can be duplicated(have the same tconst(title))**
    **because different nconst(name/person) can be included in the same tconst(title).**
  - example:
  ```
  tconst	ordering	nconst	category	job	characters
  tt0000001	1	nm1588970	self	\N	["Self"]
  tt0000001	2	nm0005690	director	\N	\N
  ```
  - characters (string) - It actually looks like "['first role', 'second role', 'etc']" or "\\N".
  - example:
  ```
  tconst	ordering	nconst	category	job	characters
  tt1834344	1	nm1129362	self	\N	["Self - Correspondent (2002-2004)","Co-Host (2004-)"]
  ```
____


### title.ratings.tsv.gz
Contains the IMDb rating and votes information for titles

  - tconst (string) - alphanumeric unique identifier of the title.
  - averageRating – weighted average of all the individual user ratings.
  - numVotes - number of votes the title has received.
____  
  
## Data cohesion
The database consist of seven files which contained tables, among which:
- The **name.basics** table contains information about the name/person and references to the titles in which the person is known (knownForTitles field).
- The tables **title.basics, title.crew, title.episode, title.ratings, title.principals** and **title.akas**,
  are linked by external ID(tconst) and contain information about the titles.
    - The table **title.basics** contains basic information about the title.
    - The table **title.crew** contains the tconst of the titles that refer to the array of directors of the table **name.basics**, that worked on them.
    - The table **title.episode** contains the tconst of the episodes, which refer to the parent ID of the title(parentTconst).
    - The table **title.principals** contains external ID(tconst) that may be duplicated,
      because the table fields refer to different name/person(nconst) of the table **name.basics**.
    - The table **title.akas** contains external ID(tconst) which may be duplicated,
      because the table fields contain different region and language of the titles.

**The relationships are illustrated in the Entity-Relationship diagram, which can be found [here](https://github.com/mk-ambivert/IMDb-searcher/master/docs/db/IMDB-ER-diagram.drawio).**

  
## IMDb dataset license details

Subsets of IMDb data are available for access to customers for personal and
non-commercial use. You can hold local copies of this data, and it is subject to
our terms and conditions. Please refer to the
[Non-Commercial Licensing](https://help.imdb.com/article/imdb/general-information/can-i-use-imdb-data-in-my-software/G5JTRESSHJBBHTGX?pf_rd_m=A2FGELUUNOQJNL&pf_rd_p=3aefe545-f8d3-4562-976a-e5eb47d1bb18&pf_rd_r=0J8FC9NDYKWB18MEW883&pf_rd_s=center-1&pf_rd_t=60601&pf_rd_i=interfaces&ref_=fea_mn_lk1#) and [copyright/license](https://www.imdb.com/conditions?pf_rd_m=A2FGELUUNOQJNL&pf_rd_p=3aefe545-f8d3-4562-976a-e5eb47d1bb18&pf_rd_r=0J8FC9NDYKWB18MEW883&pf_rd_s=center-1&pf_rd_t=60601&pf_rd_i=interfaces&ref_=fea_mn_lk2) and verify compliance.