# MARTA-arrivals

## Overview

For those who are not familiar, MARTA stands for Metropolitan Atlanta Rapid Transit Authority.

In this project, you'll be filtering data from the [MARTA Rail Realtime RESTful API](https://www.itsmarta.com/app-developer-resources.aspx).

The intent is to poll the API periodically, filter the results, and write the filtered data to a file (`out.log`). See the "Goals" section below for details.

The provided code is in [go](https://golang.org/) but feel free to start a new project from scratch if you'd prefer to use another language. Along the same lines, feel free to import additional packages / libraries as needed.

## Setup

You should be able to run the project as-is without any issues connecting to the API.

If, however, you encounter a permissions issue, feel free to [register for your own key](https://www.itsmarta.com/developer-reg-rtt.aspx) and [update the API_KEY Secret value](https://docs.replit.com/repls/secrets-environment-variables#secrets-environment-variables).

## Scenario
You've taken over this project from a summer intern who has since returned to school. It has not been deployed or even reviewed by any other developer yet. (side note: Assume that any code comments were left by the original author!)

You are tasked with updating (or rewriting in another language) as if for a real production service. (For example, assume any [stdout output will be archived](https://12factor.net/logs).)

## Goals
- Filter the original API response to only the "TRAIN_ID", "STATION", and "WAITING_SECONDS" key-value pairs. (Case-sensitivity of the resulting key/val strings is NOT important, however.)
- Implement the `filterStops` method to further filter to only items where `WAITING_SECONDS` is less than `120`.
- Importantly, to minimize disk space, this service should NOT write to `out.log` any MORE frequently than once every `10` seconds. Writing LESS-frequently, on the other hand, is acceptable, if needed, to improve resiliency.
- Audit for pre-existing bugs / easy-to-fix inefficiencies!
- Consider resiliency, idempotency, security, and error-handling.
- Handling interrupts / signals is outside the scope of this exercise, BUT -- you can assume that if the process dies, it will immediately be restarted!
- Writing tests is outside the scope of this exercise as well.
