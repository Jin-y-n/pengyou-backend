-- Active: 1722733405277@@8.137.87.209@19501@pengyou
drop database if exists pengyou_test;

create database pengyou_test;

use pengyou_test;

CREATE TABLE user
(
    id                INT UNSIGNED AUTO_INCREMENT,
    username          VARCHAR(50) NOT NULL,
    password          VARCHAR(64) NOT NULL,
    email             VARCHAR(50),
    phone             VARCHAR(20),
    login_time        TIMESTAMP   NULL     DEFAULT NULL,
    created_at        TIMESTAMP            DEFAULT CURRENT_TIMESTAMP,
    modified_at       TIMESTAMP            DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    delete_at         TIMESTAMP   NULL     DEFAULT NULL,
    status            SMALLINT             DEFAULT 1,
    heart_beat_time   TIMESTAMP   NULL     DEFAULT NULL,
    client_ip         VARCHAR(50),
    is_logout         TINYINT              DEFAULT 0,
    log_out_time      TIMESTAMP   NULL     DEFAULT NULL,
    device_info       VARCHAR(255),
    created_person    INT UNSIGNED,
    modified_person   INT UNSIGNED,
    modified_by_admin TINYINT     NOT NULL DEFAULT 0,
    PRIMARY KEY (id),
    UNIQUE INDEX idx_username (username),
    INDEX idx_email (email),
    INDEX idx_phone (phone)
);

CREATE TABLE user_profile
(
    user_id         INT UNSIGNED NOT NULL,
    display_name    VARCHAR(50),
    avatar_id       VARCHAR(255),
    bio             VARCHAR(255),
    gender          TINYINT,
    birthday        DATE,
    location        VARCHAR(100),
    occupation      VARCHAR(100),
    education       VARCHAR(100),
    school          VARCHAR(100),
    major           VARCHAR(100),
    company         VARCHAR(100),
    position        VARCHAR(100),
    website         VARCHAR(255),
    created_at      TIMESTAMP         DEFAULT CURRENT_TIMESTAMP,
    modified_at     TIMESTAMP         DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    delete_at       TIMESTAMP    NULL DEFAULT NULL,
    created_person  INT UNSIGNED,
    modified_person INT UNSIGNED,
    modified_by_admin TINYINT     NOT NULL DEFAULT 0,
    PRIMARY KEY (user_id),
    FOREIGN KEY (user_id) REFERENCES user (id) ON DELETE CASCADE
);


CREATE TABLE tag
(
    id          INT UNSIGNED AUTO_INCREMENT,
    name        VARCHAR(63) NOT NULL,
    description VARCHAR(255),
    PRIMARY KEY (id),
    UNIQUE INDEX idx_name (name)
);

CREATE TABLE user_tag_mapping
(
    id      INT UNSIGNED AUTO_INCREMENT,
    user_id INT UNSIGNED NOT NULL,
    tag_id  INT UNSIGNED NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (user_id) REFERENCES user (id) ON DELETE CASCADE,
    FOREIGN KEY (tag_id) REFERENCES tag (id) ON DELETE CASCADE
);


CREATE TABLE user_friend
(
    id             INT UNSIGNED AUTO_INCREMENT,
    user_id        INT UNSIGNED NOT NULL,
    friend_id      INT UNSIGNED NOT NULL,
    status         TINYINT           DEFAULT 0,
    request_date   TIMESTAMP    NULL DEFAULT NULL,
    accepted_date  TIMESTAMP    NULL DEFAULT NULL,
    require_person INT UNSIGNED,
    relationship   SMALLINT          DEFAULT 1,
    delete_at      TIMESTAMP    NULL DEFAULT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (user_id) REFERENCES user (id) ON DELETE CASCADE,
    FOREIGN KEY (friend_id) REFERENCES user (id) ON DELETE CASCADE,
    FOREIGN KEY (require_person) REFERENCES user (id) ON DELETE SET NULL
);


CREATE TABLE social_account
(
    id       INT UNSIGNED AUTO_INCREMENT,
    user_id  INT UNSIGNED NOT NULL,
    platform VARCHAR(63)  NOT NULL,
    link     VARCHAR(255),
    PRIMARY KEY (id),
    FOREIGN KEY (user_id) REFERENCES user (id) ON DELETE CASCADE,
    UNIQUE INDEX idx_user_platform (user_id, platform)
);



CREATE TABLE post_section
(
    id          INT UNSIGNED AUTO_INCREMENT,
    section     VARCHAR(100),
    description VARCHAR(255),
    PRIMARY KEY (id),
    UNIQUE INDEX idx_section (section)
);


CREATE TABLE post_label
(
    id          INT UNSIGNED AUTO_INCREMENT,
    label       VARCHAR(100),
    description VARCHAR(255),
    PRIMARY KEY (id),
    UNIQUE INDEX idx_label (label)
);


CREATE TABLE post
(
    id              INT UNSIGNED AUTO_INCREMENT,
    author          INT UNSIGNED NOT NULL,
    title           VARCHAR(255),
    content         TEXT,
    created_at      TIMESTAMP         DEFAULT CURRENT_TIMESTAMP,
    modified_at     TIMESTAMP         DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    status          TINYINT           DEFAULT 1,
    created_person  INT UNSIGNED,
    modified_person INT UNSIGNED,
    delete_at       TIMESTAMP    NULL DEFAULT NULL,
    modified_by_admin TINYINT     NOT NULL DEFAULT 0,

    PRIMARY KEY (id),
    FOREIGN KEY (author) REFERENCES user (id) ON DELETE CASCADE,
    FOREIGN KEY (created_person) REFERENCES user (id) ON DELETE SET NULL,
    FOREIGN KEY (modified_person) REFERENCES user (id) ON DELETE SET NULL,
    FOREIGN KEY (label) REFERENCES post_label (id) ON DELETE SET NULL
);



CREATE TABLE post_history
(
    id              INT UNSIGNED AUTO_INCREMENT,
    author          INT UNSIGNED NOT NULL,
    title           VARCHAR(255),
    content         TEXT,
    modified_at     TIMESTAMP         DEFAULT CURRENT_TIMESTAMP,
    modified_person INT UNSIGNED,
    modified_by_admin TINYINT     NOT NULL DEFAULT 0,

    delete_at       TIMESTAMP    NULL DEFAULT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (author) REFERENCES user (id) ON DELETE CASCADE,
    FOREIGN KEY (modified_person) REFERENCES user (id) ON DELETE SET NULL,
    FOREIGN KEY (label) REFERENCES post_label (id) ON DELETE SET NULL
);


CREATE TABLE post_label_mapping
(
    id       INT UNSIGNED AUTO_INCREMENT,
    post_id  INT UNSIGNED NOT NULL,
    label_id INT UNSIGNED NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (post_id) REFERENCES post (id) ON DELETE CASCADE,
    FOREIGN KEY (label_id) REFERENCES post_label (id) ON DELETE CASCADE
);


CREATE TABLE post_section_mapping
(
    id         INT UNSIGNED AUTO_INCREMENT,
    section_id INT UNSIGNED NOT NULL,
    post_id    INT UNSIGNED NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (section_id) REFERENCES post_section (id) ON DELETE CASCADE,
    FOREIGN KEY (post_id) REFERENCES post (id) ON DELETE CASCADE
);


CREATE TABLE post_history_label_mapping
(
    id              INT UNSIGNED AUTO_INCREMENT,
    post_history_id INT UNSIGNED NOT NULL,
    label_id        INT UNSIGNED NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (post_history_id) REFERENCES post_history (id) ON DELETE CASCADE,
    FOREIGN KEY (label_id) REFERENCES post_label (id) ON DELETE CASCADE
);


CREATE TABLE post_history_section_mapping
(
    id              INT UNSIGNED AUTO_INCREMENT,
    post_history_id INT UNSIGNED NOT NULL,
    section_id      INT UNSIGNED NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (post_history_id) REFERENCES post_history (id) ON DELETE CASCADE,
    FOREIGN KEY (section_id) REFERENCES post_section (id) ON DELETE CASCADE
);


CREATE TABLE post_like
(
    id         INT UNSIGNED AUTO_INCREMENT,
    post_id    INT UNSIGNED NOT NULL,
    user_id    INT UNSIGNED NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    FOREIGN KEY (post_id) REFERENCES post (id) ON DELETE CASCADE,
    FOREIGN KEY (user_id) REFERENCES user (id) ON DELETE CASCADE
);


CREATE TABLE post_dislike
(
    id         INT UNSIGNED AUTO_INCREMENT,
    post_id    INT UNSIGNED NOT NULL,
    user_id    INT UNSIGNED NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    FOREIGN KEY (post_id) REFERENCES post (id) ON DELETE CASCADE,
    FOREIGN KEY (user_id) REFERENCES user (id) ON DELETE CASCADE
);


CREATE TABLE comment
(
    id              INT UNSIGNED AUTO_INCREMENT,
    post_id         INT UNSIGNED NOT NULL,
    user_id         INT UNSIGNED NOT NULL,
    content         VARCHAR(255),
    created_time    TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    modified_time   TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    created_person  INT UNSIGNED,
    modified_person INT UNSIGNED,
    modified_by_admin TINYINT     NOT NULL DEFAULT 0,
    PRIMARY KEY (id),
    FOREIGN KEY (post_id) REFERENCES post (id) ON DELETE CASCADE,
    FOREIGN KEY (user_id) REFERENCES user (id) ON DELETE CASCADE,
    FOREIGN KEY (created_person) REFERENCES user (id) ON DELETE SET NULL,
    FOREIGN KEY (modified_person) REFERENCES user (id) ON DELETE SET NULL
);


CREATE TABLE comment_like
(
    id         INT UNSIGNED AUTO_INCREMENT,
    comment_id INT UNSIGNED NOT NULL,
    user_id    INT UNSIGNED NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    FOREIGN KEY (comment_id) REFERENCES comment (id) ON DELETE CASCADE,
    FOREIGN KEY (user_id) REFERENCES user (id) ON DELETE CASCADE
);


CREATE TABLE comment_history
(
    id              INT UNSIGNED AUTO_INCREMENT,
    post_id         INT UNSIGNED NOT NULL,
    user_id         INT UNSIGNED NOT NULL,
    content         VARCHAR(255),
    modified_time   TIMESTAMP         DEFAULT CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP,
    modified_person INT UNSIGNED,
    modified_by_admin TINYINT     NOT NULL DEFAULT 0,

    delete_at       TIMESTAMP    NULL DEFAULT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (post_id) REFERENCES post (id) ON DELETE CASCADE,
    FOREIGN KEY (user_id) REFERENCES user (id) ON DELETE CASCADE,
    FOREIGN KEY (modified_person) REFERENCES user (id) ON DELETE SET NULL
);


CREATE TABLE report
(
    id          INT UNSIGNED AUTO_INCREMENT,
    reported_id INT UNSIGNED NOT NULL,
    reporter_id INT UNSIGNED NOT NULL,
    reason      VARCHAR(255),
    report_time TIMESTAMP         DEFAULT CURRENT_TIMESTAMP,
    status      TINYINT           DEFAULT 0,
    delete_at   TIMESTAMP    NULL DEFAULT NULL,
    type        SMALLINT,
    PRIMARY KEY (id),
    FOREIGN KEY (reported_id) REFERENCES user (id) ON DELETE CASCADE,
    FOREIGN KEY (reporter_id) REFERENCES user (id) ON DELETE CASCADE
);


CREATE TABLE admin
(
    id                INT UNSIGNED AUTO_INCREMENT,
    username          VARCHAR(100) NOT NULL,
    password          VARCHAR(255) NOT NULL,
    email             VARCHAR(100),
    phone             VARCHAR(100),
    created_time      TIMESTAMP         DEFAULT CURRENT_TIMESTAMP,
    modified_time     TIMESTAMP         DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    created_person    INT UNSIGNED,
    modified_person   INT UNSIGNED      DEFAULT 0,
    delete_at         TIMESTAMP    NULL DEFAULT NULL,
    role              TINYINT           DEFAULT 0,
    modified_by_root TINYINT           DEFAULT 0,
    PRIMARY KEY (id),
    UNIQUE INDEX idx_username (username)
);


CREATE TABLE sensitive_word
(
    id   INT UNSIGNED AUTO_INCREMENT,
    word VARCHAR(100) NOT NULL,
    PRIMARY KEY (id),
    UNIQUE INDEX idx_word (word)
);


CREATE TABLE message_send
(
    id           INT UNSIGNED AUTO_INCREMENT,
    sender_id    INT UNSIGNED NOT NULL,
    recipient_id INT UNSIGNED NOT NULL,
    content      VARCHAR(511) NOT NULL,
    sent_at      TIMESTAMP         DEFAULT CURRENT_TIMESTAMP,
    is_read      TINYINT           DEFAULT 0,
    delete_at    TIMESTAMP    NULL DEFAULT NULL,
    type         TINYINT           DEFAULT 0,
    PRIMARY KEY (id),
    FOREIGN KEY (sender_id) REFERENCES user (id) ON DELETE CASCADE,
    FOREIGN KEY (recipient_id) REFERENCES user (id) ON DELETE CASCADE
);


CREATE TABLE message_receive
(
    id              INT UNSIGNED AUTO_INCREMENT,
    message_send_id INT UNSIGNED NOT NULL,
    recipient_id    INT UNSIGNED NOT NULL,
    read_at         TIMESTAMP    NULL DEFAULT NULL,
    delete_at       TIMESTAMP    NULL DEFAULT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (message_send_id) REFERENCES message_send (id) ON DELETE CASCADE,
    FOREIGN KEY (recipient_id) REFERENCES user (id) ON DELETE CASCADE
);



INSERT INTO user (id, username, password, email, phone, status)
VALUES (1, 'user1', 'hashed_password1', 'user1@example.com', '+1234567890', 1),
       (2, 'user2', 'hashed_password2', 'user2@example.com', '+1234567891', 1),
       (3, 'user3', 'hashed_password3', 'user3@example.com', '+1234567892', 1),
       (4, 'user4', 'hashed_password4', 'user4@example.com', '+1234567893', 1),
       (5, 'user5', 'hashed_password5', 'user5@example.com', '+1234567894', 1);

INSERT INTO user_profile (user_id, display_name, avatar_id, gender, birthday, location, occupation, education, school,
                          major, company, position, website)
VALUES (1, 'User One', 'avatar1', 1, '1990-01-01', 'New York', 'Developer', 'Bachelor', 'University A',
        'Computer Science', 'Tech Co.', 'Software Engineer', 'http://example.com/user1'),
       (2, 'User Two', 'avatar2', 2, '1992-02-02', 'Los Angeles', 'Designer', 'Master', 'University B',
        'Graphic Design', 'Design Inc.', 'UI Designer', 'http://example.com/user2'),
       (3, 'User Three', 'avatar3', 1, '1994-03-03', 'Chicago', 'Analyst', 'PhD', 'University C', 'Data Science',
        'Analytics Corp.', 'Data Analyst', 'http://example.com/user3'),
       (4, 'User Four', 'avatar4', 2, '1996-04-04', 'San Francisco', 'Manager', 'Bachelor', 'University D',
        'Business Administration', 'Management Ltd.', 'Project Manager', 'http://example.com/user4'),
       (5, 'User Five', 'avatar5', 1, '1998-05-05', 'Seattle', 'Engineer', 'Master', 'University E',
        'Electrical Engineering', 'Engineering Co.', 'Hardware Engineer', 'http://example.com/user5');

INSERT INTO tag (id, name, description)
VALUES (1, 'Programming', 'All about programming and coding.'),
       (2, 'Design', 'Everything related to design.'),
       (3, 'Science', 'Scientific topics and discoveries.'),
       (4, 'Travel', 'Travel experiences and destinations.'),
       (5, 'Food', 'Culinary delights and recipes.');

INSERT INTO user_tag_mapping (user_id, tag_id)
VALUES (1, 1), -- Programming
       (2, 2), -- Design
       (3, 3), -- Science
       (4, 4), -- Travel
       (5, 5); -- Food

INSERT INTO user_friend (user_id, friend_id, status)
VALUES (1, 2, 1), -- User 1 is friends with User 2
       (2, 1, 1), -- User 2 is friends with User 1
       (2, 3, 1), -- User 2 is friends with User 3
       (3, 2, 1), -- User 3 is friends with User 2
       (3, 4, 1); -- User 3 is friends with User 4

INSERT INTO social_account (user_id, platform, link)
VALUES (1, 'Twitter', 'https://twitter.com/user1'),
       (2, 'Instagram', 'https://instagram.com/user2'),
       (3, 'Facebook', 'https://facebook.com/user3'),
       (4, 'LinkedIn', 'https://linkedin.com/in/user4'),
       (5, 'GitHub', 'https://github.com/user5');

INSERT INTO post_section (id, section, description)
VALUES (1, 'Technology', 'Posts related to technology.'),
       (2, 'Lifestyle', 'Posts about lifestyle topics.'),
       (3, 'Health', 'Posts about health and wellness.'),
       (4, 'Education', 'Educational content.'),
       (5, 'Entertainment', 'Entertainment news and reviews.');

INSERT INTO post_label (id, label, description)
VALUES (1, 'Tech News', 'Latest news in technology.'),
       (2, 'DIY Projects', 'Do-it-yourself projects.'),
       (3, 'Healthy Recipes', 'Healthy cooking ideas.'),
       (4, 'Study Tips', 'Tips for studying effectively.'),
       (5, 'Movie Reviews', 'Reviews of recent movies.');

INSERT INTO post (id, author, title, content, label)
VALUES (1,1, 'My First Post', 'This is my first post on this platform.', 1),
       (2,2, 'A Day in My Life', 'Here is a glimpse into my daily routine.', 2),
       (3,3, 'Healthy Breakfast Ideas', 'Some healthy breakfast options.', 3),
       (4,4, 'How to Study Effectively', 'Tips for improving your study habits.', 4),
       (5,5, 'Top Movies of the Year', 'My top picks for the best movies this year.', 5);

INSERT INTO post_history (id, author, title, content, label)
VALUES ( 1,1, 'My First Post (History)', 'This is the original version of my first post.', 1),
       ( 2,2, 'A Day in My Life (History)', 'Here is the original version of my daily routine.', 2),
       ( 3,3, 'Healthy Breakfast Ideas (History)', 'Some original healthy breakfast options.', 3),
       ( 4,4, 'How to Study Effectively (History)', 'Original tips for improving your study habits.', 4),
       ( 5,5, 'Top Movies of the Year (History)', 'My original top picks for the best movies this year.', 5);

INSERT INTO post_label_mapping (post_id, label_id)
VALUES (1, 1), -- Tech News
       (2, 2), -- DIY Projects
       (3, 3), -- Healthy Recipes
       (4, 4), -- Study Tips
       (5, 5); -- Movie Reviews

INSERT INTO post_section_mapping (section_id, post_id)
VALUES (1, 1), -- Technology
       (2, 2), -- Lifestyle
       (3, 3), -- Health
       (4, 4), -- Education
       (5, 5); -- Entertainment

INSERT INTO post_history_label_mapping (post_history_id, label_id)
VALUES (1, 1), -- Tech News
       (2, 2), -- DIY Projects
       (3, 3), -- Healthy Recipes
       (4, 4), -- Study Tips
       (5, 5); -- Movie Reviews

INSERT INTO post_history_section_mapping (post_history_id, section_id)
VALUES (1, 1), -- Technology
       (2, 2), -- Lifestyle
       (3, 3), -- Health
       (4, 4), -- Education
       (5, 5); -- Entertainment

INSERT INTO post_like (post_id, user_id)
VALUES (1, 2),
       (2, 3),
       (3, 4),
       (4, 5),
       (5, 1);

INSERT INTO post_dislike (post_id, user_id)
VALUES (1, 3),
       (2, 4),
       (3, 5),
       (4, 1),
       (5, 2);

INSERT INTO comment (id, post_id, user_id, content)
VALUES (1,1, 2, 'Great post!'),
       (2,2, 3, 'Interesting read.'),
       (3,3, 4, 'Thanks for sharing!'),
       (4,4, 5, 'Very informative.'),
       (5,5, 1, 'Enjoyed it!');

INSERT INTO comment_like (comment_id, user_id)
VALUES (2, 4),
       (3, 5),
       (4, 1),
       (5, 2);

INSERT INTO comment_history (post_id, user_id, content)
VALUES (1, 2, 'Great post! (Original)'),
       (2, 3, 'Interesting read. (Original)'),
       (3, 4, 'Thanks for sharing! (Original)'),
       (4, 5, 'Very informative. (Original)'),
       (5, 1, 'Enjoyed it! (Original)');

INSERT INTO report (reported_id, reporter_id, reason)
VALUES (1, 2, 'Inappropriate content'),
       (2, 3, 'Spamming'),
       (3, 4, 'Misleading information'),
       (4, 5, 'Hateful speech'),
       (5, 1, 'Breach of privacy');

INSERT INTO admin (id, username, password, email, phone, role)
VALUES (1, 'admin1', 'hashed_password1', 'admin1@example.com', '+1234567890', 1),
       (2, 'admin2', 'hashed_password2', 'admin2@example.com', '+1234567891', 1),
       (3, 'admin3', 'hashed_password3', 'admin3@example.com', '+1234567892', 1),
       (4, 'admin4', 'hashed_password4', 'admin4@example.com', '+1234567893', 1),
       (5, 'admin5', 'hashed_password5', 'admin5@example.com', '+1234567894', 1);


INSERT INTO sensitive_word (word)
VALUES ('offensive'),
       ('hate'),
       ('spam'),
       ('virus'),
       ('malware');

INSERT INTO message_send (id, sender_id, recipient_id, content)
VALUES (1,1, 2, 'Hello there!'),
       (2,2, 3, 'How are you?'),
       (3,3, 4, 'Nice to meet you.'),
       (4,4, 5, 'Looking forward to your next post.'),
       (5,5, 1, 'Thank you for your support.');

INSERT INTO message_receive (id, message_send_id, recipient_id)
VALUES (1,1, 2),
       (2,2, 3),
       (3,3, 4),
       (4,4, 5),
       (5,5, 1);

