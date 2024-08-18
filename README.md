# Modular Social App

PengYou is an application designed to attract people with different interests to communicate, chat, and become friends through various communities. Our goal is to help users find compatible partners.

- [Frontend Project](https://github.com/Napbad/pengyou-frontend)

## Project Requirements Analysis

### Core Functional Requirements

#### User Module

- **Permissions**
  - **Registration**
    - Support registration via email or mobile number, with email verification through IMAP/POP protocols and SMS verification integrated with Alibaba Cloud services.
    - Both email and mobile numbers must undergo dual-factor verification to ensure security.
  - **Login**
    - Support multiple login methods: username, email, mobile number with password or verification code.
    - Upon successful login, the system returns a JWT token for subsequent request authentication.
  - **Logout**
    - Clear session and locally stored JWT tokens to ensure user privacy.
  - **Password Recovery**
    - Send verification codes via mobile or email to allow users to reset their passwords, requiring dual-factor verification.

- **Profile**
  - Users can edit personal information, including but not limited to:
    - Personality type
    - Username
    - Avatar
    - Bio
    - Gender
    - Birthday
    - Location
    - Occupation
    - Job title
    - Education background
    - School
    - Major
    - Tags

- **Chatting**
  - Support one-to-one chats with message types including text, images, videos, GIFs, and file transfers.
  - Chat history is saved, allowing users to view past conversations at any time.
  - Users can add, delete friends, and search for friends by keyword.

#### Community Module

- **Posting**
  - Users can post videos and text content to the community.
  - Content is reviewed before posting (automatically detects sensitive words) to ensure compliance with community guidelines.

- **Comments**
  - Users can comment on posts with support for multi-level replies.
  - Comments are reviewed before posting (automatically detects sensitive words) to ensure legality and compliance.
  - Comment management, including reporting and removing inappropriate comments.

- **Categories**
  - The community is divided into thematic areas, allowing users to browse and participate according to their interests.

- **Content Management**
  - Users can edit and delete their own posts and comments.

#### Search

- Implement keyword and feature-based search functionality for users, posts, communities, and comments.
- Search results are sorted by relevance, with filtering and sorting options provided.

#### Backend Management

- **User Management**
  - View, edit, and disable user accounts.
  - Monitor user behavior and handle violations.
- **Content Review**
  - Review user posts and comments to ensure compliance.
  - Accept and process reports.
  - Manually review controversial content (when automatic sensitive word detection fails).
- **System Monitoring**
  - Monitor system performance and promptly address issues.

### Optional Functional Requirements

- **Friend Groups**
- **Groups**
  - Join, leave groups, create and manage group information (admin), invite members.
- **Report Inappropriate Content**
- **System Matching and Recommendations**
- **Follow**
- **Instant Message Notifications**
- **Comment Notifications**
- **Dynamic Update Notifications**
- **Third-party Login**
  - WeChat, QQ, etc.
- **User Verification**
- **Nearby**
- **Instant Casual Chat**
- **Instant Casual Games**
- **AI Integration**
  - Chatbots, intelligent recommendations, etc.

---

## Technology Stack

- **Backend**: Go (or another backend technology)
- **Frontend**: Vue.js (or another frontend technology)
- **Database**: MySQL / MongoDB (or another database technology)
- **Other**: JWT, Alibaba Cloud Services, automatic sensitive word detection, etc.

---

## Development Environment Setup



```bash
cd ./install/dev
```

* init
    * if you are **Linux** user

      run
      ```bash
      ./linux/init.sh
      ```

    * if you are **Windows** user

      run
      ```bash
      ./windows/init-pw.sh
      ```
      in powershell

* database

after that, you need to run init.sql in your mysql container

like this:
```bash
podman exec -it mysql1 mysql -u root -p 12345678 
```
    Then, enter the content of `init.sql`.

## Contact Us

- For any questions or suggestions, please contact us via:
    - Email: napbad.sen@gmail.com
    - GitHub Issues: [Click here](https://github.com/Napbad/pengyou-backend/issues/new)

---

## Copyright Statement

- This project is licensed under the Apache License. More information can be found in the [LICENSE](LICENSE) file.
