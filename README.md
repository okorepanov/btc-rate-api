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

# Endpoints results

<img width="970" alt="Screenshot 2022-07-29 at 12 10 33" src="https://user-images.githubusercontent.com/59099358/181728006-1b7d26cc-b90e-48d6-94f1-448961a2c38b.png">
<img width="963" alt="Screenshot 2022-07-29 at 12 11 33" src="https://user-images.githubusercontent.com/59099358/181728018-cf9a95fc-7aeb-4033-9ba1-31a4d051e225.png">
<img width="974" alt="Screenshot 2022-07-29 at 12 24 11" src="https://user-images.githubusercontent.com/59099358/181728871-e425efe9-7af2-48e1-b578-a64f724361ea.png">


# Useful Links
App password generation: https://www.lifewire.com/get-a-password-to-access-gmail-by-pop-imap-2-1171882
