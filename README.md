# Country Information Service API

This is a REST web application built with Go that provides information about countries, including general info, historical population data, and a diagnostic status overview. The service is designed to interact with the following third-party APIs:
- [CountriesNow API](https://documenter.getpostman.com/view/1134062/T1LJjU52)
- [REST Countries API](http://129.241.150.113:8080/)

The service supports three main endpoints:
1. General Country Info
2. Historical Population Data
3. Service Diagnostics

---

## Endpoints

### 1. Country Info Endpoint
Returns general information for a given country using its 2-letter ISO code.

- **Method**: `GET`
- **Path**: `/countryinfo/v1/info/{:two_letter_country_code}{?limit=10}`
- **Example**:
    ```http
    GET https://countryinfoservice-8oob.onrender.com/countryinfo/v1/info/no
    GET https://countryinfoservice-8oob.onrender.com/countryinfo/v1/info/no?limit=5
    ```
- **Query Parameters**:
    - `two_letter_country_code` (required): ISO 3166-2 code (e.g., `no` for Norway).
    - `limit` (optional): Limits the number of cities listed in ascending alphabetical order.

- **Response (200 OK)**:
    ```json
    {
      "name": "Norway",
      "continents": ["Europe"],
      "population": 4700000,
      "languages": {"nno":"Norwegian Nynorsk","nob":"Norwegian Bokmål","smi":"Sami"},
      "borders": ["FIN","SWE","RUS"],
      "flag": "https://flagcdn.com/w320/no.png",
      "capital": "Oslo",
      "cities": ["Abelvaer","Adalsbruk","Adland"]
    }
    ```

---

### 2. Country Population Endpoint
Returns historical population data for a given country and calculates the mean value.

- **Method**: `GET`
- **Path**: `/countryinfo/v1/population/{:two_letter_country_code}{?limit={:startYear-endYear}}`
- **Example**:
    ```http
    GET https://countryinfoservice-8oob.onrender.com/countryinfo/v1/population/no
    GET https://countryinfoservice-8oob.onrender.com/countryinfo/v1/population/no?limit=2010-2015
    ```
- **Query Parameters**:
    - `two_letter_country_code` (required): ISO 3166-2 code.
    - `limit` (optional): Limits the historical data to the given year range.

- **Response (200 OK)**:
    ```json
    {
      "mean": 5044396,
      "values": [
          {"year":2010,"value":4889252},
          {"year":2011,"value":4953088},
          {"year":2012,"value":5018573},
          {"year":2013,"value":5079623},
          {"year":2014,"value":5137232},
          {"year":2015,"value":5188607}
      ]
    }
    ```

---

### 3. Diagnostics Endpoint
Provides a status overview of the services this API depends on.

- **Method**: `GET`
- **Path**: `/countryinfo/v1/status/`
- **Example**:
    ```http
    GET http://<base_url>/countryinfo/v1/status/
    ```
- **Response (200 OK)**:
    ```json
    {
      "countriesnowapi": 200,
      "restcountriesapi": 200,
      "version": "v1",
      "uptime": 123456
    }
    ```
    - `countriesnowapi`: HTTP status of the CountriesNow API.
    - `restcountriesapi`: HTTP status of the REST Countries API.
    - `version`: Version of this API.
    - `uptime`: Time in seconds since last restart.

---


## Deployment
This service is deployed on Render. You can access the live version at: https://countryinfoservice-8oob.onrender.com 
