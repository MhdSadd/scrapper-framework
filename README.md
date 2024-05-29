# Scraper Framework

**main** ---------------> **config manager** --------------> **main** --------------> **scheduler** --------------> **extractor** --------------> **scheduler** --------------> **main** ?--------------> **transformer** --------------> **main**

The diagram above shows the flow of data/control from the start of the scraping process to the end of it.
The details of what happens at each point is captured below in the [`Scraping Process Explained`](#scraping-process-explained).

### Scraping Process Explained

> **main:**
>   - Calls the **config manager** to get the config file path and validate
>
> **config manager:**
>   - Gets path to config file via terminal
>   - Validates the config file
>   - Returns data to the **main**
>
> **main:** 
>   - Gets the response data from the **config manager**
>   - Sends the data to the **scheduler**, if the config file was valid.
>   - If the file is invalid, tells the user that the config file is invalid.
>
> **scheduler:**
>    - Schedules requests to the **extractor**
>
>
> **extractor:**
>    - Extracts the info from the website and save it to a file based on the structure specified in the entity
>    - Returns the extracted data to the **scheduler**
>
> **scheduler:**
>    - Returns the data gotten from the **extractor** to the main
>
> **main:**
>    - Calls the **transformer** if the config file provived by the user has a **transformer**
>
>
> **transformer:**
>    - Transforms the data and save it to file
>    - Returns success value to **main**
>
> **main:**
>    - Lets the user know that the scraping process is done
>    - Lets the user know if Transformation was successful