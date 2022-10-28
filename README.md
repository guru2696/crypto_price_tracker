# GoLang Assignment

**Requirements**

Write a `crypto price tracker`, a service that calls  `Coin Desk API` fetches data, parses it, and provides the
below-defined response to the user. Also, store it in internal memory using a background process.

If data available in memory, is within `expiry` then return data from memory instead of calling the API. `Expiry` should
be configurable using the environment variable.

- Log request and response using middleware.
- Ensure design is modular (Should be able to support other API providers)
- Use go language and any packages you like.
- Use gin framework [https://github.com/gin-gonic/gin].
- Follow API design
  guidelines [[Web-design-the-missing-link-ebook-2016-11.pdf (apigee.com)](https://docs.apigee.com/files/Web-design-the-missing-link-ebook-2016-11.pdf)]

**Process to follow**

- Create linear ticket in notion if not already created.
- Daily, add your updates to linear ticket
    - updates
    - todo
    - blocker

  should be separately mentioned on linear ticket before standup.

- Discuss with the assignment provider, about various approaches.
- This assignment should be considered as a first actual task.

**How to submit assignment?**

- Create private repo in your personal github account.
- Create develop and master branches.
- Push your code to `develop` branch and raise PR to `master` branch.
- Invite reviewer to the repo and assign for review.

**Expected timeline**

1 week

**Data**
Coin Desk API - [`https://api.coindesk.com/v1/bpi/currentprice.json`](https://api.coindesk.com/v1/bpi/currentprice.json)
Expected response structure -

```jsx
{
    "data": {
        "bitcoin": {
            "EUR": "43,947.8947", 
            "USD": "49,822.2917"
        }
    }
}
```

**Extended Requirements -**

- Support [cryptonator](https://api.cryptonator.com/api/ticker/btc-usd) API provider.
- Introduce sql or nosql DB for storing data.