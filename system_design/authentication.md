# Authentication

<br>
<br>
<br>
<br>
<br>

## Index

- [Session](#session)
    - [How session works](#how-session-works)
    - [Advantages](#advantages)
    - [Disadvantages](#disadvantages)
    - [Common session attacks](#common-session-attacks)
    - [Sticky Sessions](#sticky-sessions)
- [Cookie](#cookie)
    - [Types of cookies](#types-of-cookies)
    - [Important cookie attributes](#important-cookie-attributes)
- [OAuth / OA](#oauth--oauth-20)
    - [OAuth vs OAuth 2.0](#oauth-vs-oauth-20)
    - [How OAuth 2.0 works?](#how-oauth-20-works)
    - [Advantages](#advantages-1)
    - [Disadvantages](#disadvantages-1)
    - [OpenID Connect (OIDC)](#openid-connect-oidc)

<br>
<br>
<br>
<br>
<br>

### Session

- A session is a server-side mechanism to remember a user across multiple HTTP requests.
- HTTP is stateless and every request is independent. A session adds state so the server can recognize who is making the request.

<br>

#### How session works

1. A user tries to login with a username and password.
2. The server verifies the credentials. If the credentials are corrrect, the user is logged in and a session is generated for that user.
3. Each session contains data like `session_id`, `permissions`, `login_time`, etc. This session is stored on the server side.
4. There are multiple options for storing the session:
    - Web server memory. Used with **Sticky Sessions**.
    - Distributed caches like Redis & Memcached. Used in distributed systems.
    - Database. Used when we want persistent sessions.
4. After the session is stored, its `session_id` is sent to the client, usually via a `cookie`.
```py
# response headers
Set-Cookie: session_id=abc123; HttpOnly; Secure
```
5. Client's brower will then save this `cookie` and sent it with every following request to the server.
6. Server can then make use of the `session_id` present inside this `cookie` to identify the user making the request.
7. When the user logs out, we will remove or expire the session from the server side. It is also a good practise to remove the `session_id` on the client side.
```py
# response headers
Set-Cookie: session_id=; Max-Age=0; HttpOnly; Secure
```

<br>

#### Advantages

- Security: Sensitive data never goes to the client. Client only stores a random ID.
- Easy to Revoke: To revoke a session, just remove the session from the server side. The user is instantly logged out.
- Flexibility: Can store roles, permissions and other required data in sessions.

<br>

#### Disadvantages

- Scalability:
    - More the number of users, more data we have to store. This makes scaling expensive & costly.
    - In a distributes system, we need a common session storage that will be accessed by all the servers.
- Not Ideal for APIs / Mobile:
    - Cookies don't work well for native apps.
    - Harder to manage secure sessions for cross-domain APIs.

<br>

#### Common session attacks

Session Hijacking

- An attacker steals a valid session identifier (cookie / token) and impersonates a user.
- Common causes:
    - Session IDs sent over HTTP.
    - Session ID in URLs or logs.
    - XSS exposing cookies.
- Mitigations:
    - Enforce HTTPS everywhere.
    - Never put session IDs in URL params & console logs.
    - Use these attributes when setting a cookie:
        - `Secure`: Cookie sent only over HTTPS. Prevents MITM attacks.
        - `HttpOnly`: Cookie cannot be accessed by JavaScript. Protects against XSS.
        - `SameSite=Lax / Strict`: Cookie sent on top-level navigation. Reduce CSRF attacks.

---

Session Fixation

- Victim validates the session ID already known and set by attackers.
- How it works:
    - Victim clicks on a link with a predefined `session_id` and performs login.
    - The server reuses the `session_id` present in the URL.
    - Once validated, attackers uses this `session_id` to make requests to the server.
- Mitigations:
    - Always regenerate session ID on authentication.
    - Ignore session IDs provided in URLs & request body.

<br>

#### Sticky Sessions

- A sticky session (also called **Session Affinity**) is a technique where a user's requests are consistently routed to the same backend server during a session.
- When an application stores session state on the server, the load balancer needs a way to ensure subsequent requests from the same user reach the same server.

---

How it works?

- A load balancer keeps track of which server handled the first request and routes future requests from that user to the same server.
- Load balancer inserts a cookie like `SERVER_ID=server-3` in the response of the first request. The browser then sends it back on each request.

<br>
<br>
<br>

### Cookie

- A cookie is a small piece of data stored in the client's browser.
- It is automatically sent to the server with every HTTP request to the same domain.
- Note that the frontend code does not need to manually attach cookies.
- Cookies have a size limits, usually around 4 KB per cookie.

<br>

#### Types of cookies

Session Cookies

- Temporary cookies stored in memory and deleted when the browser closes.
- Used for things like keeping you logged in during a session.

Persistent Cookies

- Stored on the user’s device for a specific period (set by the `Expires` or `Max-Age` attribute).
- Used for remembering preferences, login info, or tracking behavior over time.

<br>

#### Important cookie attributes:

`HttpOnly`

- Prevents client-side scripts (like JavaScript) from accessing the cookie.
- Helps protect against cross-site scripting (XSS) attacks that try to steal session cookies.

`Secure`

- Ensures the cookie is only sent over HTTPS connections.
- Protects cookies from being intercepted over unencrypted (HTTP) networks.
- Prevents Man-in-the-Middle (MITM) attacks.

`SameSite`

- Controls whether cookies are sent with cross-site requests.
- Helps mitigate cross-site request forgery (CSRF) by limiting when cookies are included.
- Valid values:
    - `Strict`: The cookie is only sent for requests originating from the same site.
    - `Lax`: The cookie is sent with same-site requests and some safe cross-site requests, like top-level navigation GET requests (e.g., following a link).
    - `None`: The cookie is sent with all requests, including cross-site requests.

`Expires`

- Specifies the exact date and time when the cookie should expire.
- After this date, the browser automatically deletes the cookie.

`Max-Age`

- Specifies the lifetime of the cookie in seconds from the moment it’s set.
- Overrides Expires if both are present, controlling how long the cookie persists.

<br>
<br>
<br>

### OAuth / OAuth 2.0

- OAuth stands for **Open Authorization**.
- It is an authorization framework that allows a third-party application to access a user's resources without sharing the user's credentials (password).
- OAuth **does NOT** handle authentication by itself, it only handles authorization.
- Example: Allow Spotify to read my Google contacts without sharing my Google password.

<br>

#### OAuth vs OAuth 2.0

- OAuth is the original authorization framework (v1.0) with complex cryptographic signing.
- OAuth 2.0 is a redesigned, simpler, and more flexible version that uses tokens and is widely adopted today.

<br>

#### How OAuth 2.0 works?

Actors:

| Role                     | Description                                  |
| ------------------------ | -------------------------------------------- |
| **Resource Owner**       | The user (you)                               |
| **Client**               | App requesting access (Spotify)              |
| **Authorization Server** | Issues tokens (Google Auth)                  |
| **Resource Server**      | API holding data (Google Contacts API)       |

<br>

High level flow:

- User clicks on "Login with Google".
- He is redirected to Google's **Authorization Server**. He enters his email & password and logs in.
- After login, auth server asks for user consent to share the required information.
- If the user accepts, it redirects back to the client app with an authorization code.
- The client exchanges this authorization code for access (and possibly refresh) tokens.
- Now, when the **Client** app needs to access contacts on behalf of the user, it makes request to the Google Contacts API using the access token.

---

- When the user logs out of the client app, the client will simply delete the stored access and refresh token.
- Optionally, refresh tokens can be revoked by calling the provider's revocation endpoint (recommended for security).

<br>

#### Advantages

- Security: No password sharing with the client app.
- Access Control: The issued tokens have limited access to the resource servers.
- Scalability: JWT tokens are stateless.

<br>

#### Disadvantages

- OAuth is an authorization framework, it requires OIDC for authentication.
- Token management overhead: Rotation, Revocation, Expiry handling.
- Complexity: Has many flows and configurations.

<br>

#### OpenID Connect (OIDC)

- OIDC stands for OpenID Connect.
- It is an authentication layer built on top of OAuth 2.0.

High level flow:

- User clicks on "Login with Google".
- He is redirected to Google's **Authorization Server**. He enters his email & password and logs in.
- After login, user is redirected back to the client app with an authorization code.
- Client app then sends the authorization code to Google's token endpoint.
- This endpoint then returns:
    - ID Token (JWT → proves user identity)
    - Access Token (to call APIs)
    - (Optional) Refresh Token

<br>
<br>
<br>

### 

<br>
<br>
<br>
