
# Twilio Go Verify: 2FA with Golang & Vue.js

A reference implementation for building secure multi-factor authentication (MFA) in Go using [Twilio Verify](https://www.twilio.com/verify) (Authy & SMS OTP) and a Vue.js SPA frontend.

## Features

- User registration and login with password
- SMS OTP verification using Twilio Verify
- TOTP (Authy/Authenticator app) setup and verification
- Simple in-memory user store (for demo)
- Vue.js SPA for user flows (signup, login, SMS, TOTP, dashboard)
- Clean separation of backend (Go) and frontend (Vue.js)

## Architecture

```
Golang (API, Twilio integration)
│
├── /cmd/main.go         # Server entrypoint
├── /internal
│   ├── /handlers        # HTTP handlers (auth, verify, user)
│   ├── /models          # User model
│   ├── /store           # In-memory user store
│   └── /twilio          # Twilio Verify client logic
└── /static              # Vue.js SPA (HTML, JS, CSS)
```

## Prerequisites

- Go 1.23+
- Node.js (for Vue.js, optional for static demo)
- Twilio account with Verify service ([get started](https://www.twilio.com/verify))
- Twilio credentials: `TWILIO_ACCOUNT_SID`, `TWILIO_AUTH_TOKEN`, `TWILIO_VERIFY_SID`

## Setup

1. **Clone the repo**
	```bash
	git clone https://github.com/DesmondSanctity/twilio-go-verify.git
	cd twilio-go-verify
	```

2. **Configure environment**
	- Copy `.env.example` to `.env` and fill in your Twilio credentials:
	  ```
	  TWILIO_ACCOUNT_SID=ACxxxxxxxxxxxxxxxxxxxx
	  TWILIO_AUTH_TOKEN=your_auth_token
	  TWILIO_VERIFY_SID=VAxxxxxxxxxxxxxxxxxxxx
	  ```

3. **Install Go dependencies**
	```bash
	go mod tidy
	```

4. **Run the server**
	```bash
	go run ./cmd/main.go
	```
	The server runs on [http://localhost:8080](http://localhost:8080).

5. **Access the frontend**
	- Open your browser to [http://localhost:8080](http://localhost:8080).
	- The Vue.js SPA is served from `/static`.

## Usage Flow

1. **Sign Up:** Register with name, email, password, and phone number.
2. **Login:** Enter email and password.
3. **SMS Verification:** Receive and enter OTP sent via SMS.
4. **TOTP Setup:** Scan QR code with Authy/Google Authenticator, enter generated code.
5. **Dashboard:** Access protected area after successful 2FA.

## File Structure

- **Backend (Go):**
  - `cmd/main.go`: Server setup, routing, static file serving.
  - `internal/handlers/`: Auth, user, and verification endpoints.
  - `internal/twilio/`: Twilio Verify API integration.
  - `internal/store/`: In-memory user storage.
  - `internal/models/`: User struct.

- **Frontend (Vue.js SPA in `/static/js/components/`):**
  - `Signup.js`, `Login.js`: Auth forms.
  - `SMSVerification.js`: SMS OTP input.
  - `AuthySetup.js`: TOTP QR code and code input.
  - `Dashboard.js`: Post-2FA landing page.

## API Endpoints (Sample)

- `POST /api/signup` — Register user
- `POST /api/login` — Login user
- `POST /api/verify/send-sms` — Send SMS OTP
- `POST /api/verify/verify-sms` — Verify SMS OTP
- `POST /api/verify/setup-totp` — Generate TOTP QR
- `POST /api/verify/verify-totp` — Verify TOTP code

## Security Notes

- This demo uses an in-memory store (not for production).
- Passwords are not hashed (add hashing for real use).
- Add HTTPS, CORS, and session management for production.

## References

- [Twilio Verify Docs](https://www.twilio.com/docs/verify)
- [Twilio Blog: Multi-factor Auth in Go](https://www.twilio.com/en-us/blog/developers/community/multi-factor-authentication-go-twilio-verify)
- [Vue.js](https://vuejs.org/)

## License

MIT

