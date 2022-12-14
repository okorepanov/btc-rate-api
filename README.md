# btc-rate-api
GO API project, that works with BTC-UAH rate

# Endpoints
1. **GET** /rate - returns current BTC-UAH rate
2. **POST** /subscription - saves email (that passed thorough request body) to subscribers
3. **POST** /sendEmail - sends current BTC-UAH rate to subscribers

# Setup
After cloning this repo, there needed few more steps before working with it:
1. For successful email sending add your mail credentials to ```.env``` file (EMAIL_ADDRESS, PASSWORD).
>If you have 2-factor authentication on your account, you have to pass your application-password to PASSWORD field.

2. To work via Dockerfile yo have to:
  - run ```docker build --tag btc-rate-api .```
  - run ```docker run -d -p 8080:8080 btc-rate-api```
> App will response for http://localhost:8080 url

# Endpoints statuses

1. **GET** /rate
  - 200 - returns BTC-UAH rate
2. **POST** /subscription
  - 201 - successfully save email
  - 400 - email is not valid
  - 409 - emails is subcribed already
3. **POST** /sendEmail
  - 200 - messages were successfully sent

# Endpoints results

<img width="970" alt="Screenshot 2022-07-29 at 12 10 33" src="https://user-images.githubusercontent.com/59099358/181728006-1b7d26cc-b90e-48d6-94f1-448961a2c38b.png">
<img width="975" alt="Screenshot 2022-07-30 at 15 12 24" src="https://user-images.githubusercontent.com/59099358/181913711-10703c74-42f7-4ab6-91e2-a1b057f8a46b.png">
<img width="974" alt="Screenshot 2022-07-29 at 12 24 11" src="https://user-images.githubusercontent.com/59099358/181728871-e425efe9-7af2-48e1-b578-a64f724361ea.png">

# Useful Links
App password generation: https://www.lifewire.com/get-a-password-to-access-gmail-by-pop-imap-2-1171882
