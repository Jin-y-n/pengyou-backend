-- Active: 1722733405277@@8.137.87.209@19501@pengyou
drop database if exists pengyou_test;

create database pengyou_test;

use pengyou_test;

create table admin
(
    id               int unsigned auto_increment
        primary key,
    username         varchar(100)                           not null,
    password         varchar(255)                           not null,
    email            varchar(100)                           null,
    phone            varchar(100)                           null,
    created_time     timestamp    default CURRENT_TIMESTAMP null,
    modified_time    timestamp    default CURRENT_TIMESTAMP null on update CURRENT_TIMESTAMP,
    created_person   int unsigned                           null,
    modified_person  int unsigned default '0'               null,
    delete_at        timestamp                              null,
    role             tinyint      default 0                 null,
    modified_by_root tinyint      default 0                 null,
    constraint idx_username
        unique (username)
);

create table post_label
(
    id          int unsigned auto_increment
        primary key,
    label       varchar(100) null,
    description varchar(255) null,
    constraint idx_label
        unique (label)
);

create table post_section
(
    id          int unsigned auto_increment
        primary key,
    section     varchar(100) null,
    description varchar(255) null,
    constraint idx_section
        unique (section)
);

create table sensitive_word
(
    id   int unsigned auto_increment
        primary key,
    word varchar(100) not null,
    constraint idx_word
        unique (word)
);

create table tag
(
    id          int unsigned auto_increment
        primary key,
    name        varchar(63)  not null,
    description varchar(255) null,
    constraint idx_name
        unique (name)
);

create table user
(
    id                int unsigned auto_increment
        primary key,
    username          varchar(50)                         not null,
    password          varchar(64)                         not null,
    email             varchar(50)                         null,
    phone             varchar(20)                         null,
    login_time        timestamp                           null,
    created_at        timestamp default CURRENT_TIMESTAMP null,
    modified_at       timestamp default CURRENT_TIMESTAMP null on update CURRENT_TIMESTAMP,
    delete_at         timestamp                           null,
    status            smallint  default 1                 null,
    heart_beat_time   timestamp                           null,
    client_ip         varchar(50)                         null,
    is_logout         tinyint   default 0                 null,
    log_out_time      timestamp                           null,
    device_info       varchar(255)                        null,
    created_person    int unsigned                        null,
    modified_person   int unsigned                        null,
    modified_by_admin tinyint   default 0                 not null,
    constraint idx_username
        unique (username)
);

create table message_send
(
    id           int unsigned auto_increment
        primary key,
    sender_id    int unsigned                        not null,
    recipient_id int unsigned                        not null,
    content      varchar(511)                        not null,
    sent_at      timestamp default CURRENT_TIMESTAMP null,
    is_read      tinyint   default 0                 null,
    delete_at    timestamp                           null,
    type         tinyint   default 0                 null,
    constraint message_send_ibfk_1
        foreign key (sender_id) references user (id)
            on delete cascade,
    constraint message_send_ibfk_2
        foreign key (recipient_id) references user (id)
            on delete cascade
);

create table message_receive
(
    id              int unsigned auto_increment
        primary key,
    message_send_id int unsigned not null,
    recipient_id    int unsigned not null,
    read_at         timestamp    null,
    delete_at       timestamp    null,
    constraint message_receive_ibfk_1
        foreign key (message_send_id) references message_send (id)
            on delete cascade,
    constraint message_receive_ibfk_2
        foreign key (recipient_id) references user (id)
            on delete cascade
);

create index message_send_id
    on message_receive (message_send_id);

create index recipient_id
    on message_receive (recipient_id);

create index recipient_id
    on message_send (recipient_id);

create index sender_id
    on message_send (sender_id);

create table post
(
    id                int unsigned auto_increment
        primary key,
    author            int unsigned                        not null,
    title             varchar(255)                        null,
    content           text                                null,
    created_at        timestamp default CURRENT_TIMESTAMP null,
    modified_at       timestamp default CURRENT_TIMESTAMP null on update CURRENT_TIMESTAMP,
    status            tinyint   default 1                 null,
    created_person    int unsigned                        null,
    modified_person   int unsigned                        null,
    delete_at         timestamp                           null,
    modified_by_admin tinyint   default 0                 not null,
    constraint post_ibfk_1
        foreign key (author) references user (id)
            on delete cascade,
    constraint post_ibfk_2
        foreign key (created_person) references user (id)
            on delete set null,
    constraint post_ibfk_3
        foreign key (modified_person) references user (id)
            on delete set null
);

create table comment
(
    id                int unsigned auto_increment
        primary key,
    post_id           int unsigned                        not null,
    user_id           int unsigned                        not null,
    content           varchar(255)                        null,
    created_time      timestamp default CURRENT_TIMESTAMP null,
    modified_time     timestamp default CURRENT_TIMESTAMP null on update CURRENT_TIMESTAMP,
    created_person    int unsigned                        null,
    modified_person   int unsigned                        null,
    modified_by_admin tinyint   default 0                 not null,
    constraint comment_ibfk_1
        foreign key (post_id) references post (id)
            on delete cascade,
    constraint comment_ibfk_2
        foreign key (user_id) references user (id)
            on delete cascade,
    constraint comment_ibfk_3
        foreign key (created_person) references user (id)
            on delete set null,
    constraint comment_ibfk_4
        foreign key (modified_person) references user (id)
            on delete set null
);

create index created_person
    on comment (created_person);

create index modified_person
    on comment (modified_person);

create index post_id
    on comment (post_id);

create index user_id
    on comment (user_id);

create table comment_history
(
    id                int unsigned auto_increment
        primary key,
    post_id           int unsigned                        not null,
    user_id           int unsigned                        not null,
    content           varchar(255)                        null,
    modified_time     timestamp default CURRENT_TIMESTAMP null on update CURRENT_TIMESTAMP,
    modified_person   int unsigned                        null,
    modified_by_admin tinyint   default 0                 not null,
    delete_at         timestamp                           null,
    constraint comment_history_ibfk_1
        foreign key (post_id) references post (id)
            on delete cascade,
    constraint comment_history_ibfk_2
        foreign key (user_id) references user (id)
            on delete cascade,
    constraint comment_history_ibfk_3
        foreign key (modified_person) references user (id)
            on delete set null
);

create index modified_person
    on comment_history (modified_person);

create index post_id
    on comment_history (post_id);

create index user_id
    on comment_history (user_id);

create table comment_like
(
    id         int unsigned auto_increment
        primary key,
    comment_id int unsigned                        not null,
    user_id    int unsigned                        not null,
    created_at timestamp default CURRENT_TIMESTAMP null,
    constraint comment_like_ibfk_1
        foreign key (comment_id) references comment (id)
            on delete cascade,
    constraint comment_like_ibfk_2
        foreign key (user_id) references user (id)
            on delete cascade
);

create index comment_id
    on comment_like (comment_id);

create index user_id
    on comment_like (user_id);

create index author
    on post (author);

create index created_person
    on post (created_person);

create index modified_person
    on post (modified_person);

create table post_dislike
(
    id         int unsigned auto_increment
        primary key,
    post_id    int unsigned                        not null,
    user_id    int unsigned                        not null,
    created_at timestamp default CURRENT_TIMESTAMP null,
    constraint post_dislike_ibfk_1
        foreign key (post_id) references post (id)
            on delete cascade,
    constraint post_dislike_ibfk_2
        foreign key (user_id) references user (id)
            on delete cascade
);

create index post_id
    on post_dislike (post_id);

create index user_id
    on post_dislike (user_id);

create table post_history
(
    id                int unsigned auto_increment
        primary key,
    author            int unsigned                        not null,
    title             varchar(255)                        null,
    content           text                                null,
    modified_at       timestamp default CURRENT_TIMESTAMP null,
    modified_person   int unsigned                        null,
    modified_by_admin tinyint   default 0                 not null,
    delete_at         timestamp                           null,
    constraint post_history_ibfk_1
        foreign key (author) references user (id)
            on delete cascade,
    constraint post_history_ibfk_2
        foreign key (modified_person) references user (id)
            on delete set null
);

create index author
    on post_history (author);

create index modified_person
    on post_history (modified_person);

create table post_history_label_mapping
(
    id              int unsigned auto_increment
        primary key,
    post_history_id int unsigned not null,
    label_id        int unsigned not null,
    constraint post_history_label_mapping_ibfk_1
        foreign key (post_history_id) references post_history (id)
            on delete cascade,
    constraint post_history_label_mapping_ibfk_2
        foreign key (label_id) references post_label (id)
            on delete cascade
);

create index label_id
    on post_history_label_mapping (label_id);

create index post_history_id
    on post_history_label_mapping (post_history_id);

create table post_history_section_mapping
(
    id              int unsigned auto_increment
        primary key,
    post_history_id int unsigned not null,
    section_id      int unsigned not null,
    constraint post_history_section_mapping_ibfk_1
        foreign key (post_history_id) references post_history (id)
            on delete cascade,
    constraint post_history_section_mapping_ibfk_2
        foreign key (section_id) references post_section (id)
            on delete cascade
);

create index post_history_id
    on post_history_section_mapping (post_history_id);

create index section_id
    on post_history_section_mapping (section_id);

create table post_label_mapping
(
    id       int unsigned auto_increment
        primary key,
    post_id  int unsigned not null,
    label_id int unsigned not null,
    constraint post_label_mapping_ibfk_1
        foreign key (post_id) references post (id)
            on delete cascade,
    constraint post_label_mapping_ibfk_2
        foreign key (label_id) references post_label (id)
            on delete cascade
);

create index label_id
    on post_label_mapping (label_id);

create index post_id
    on post_label_mapping (post_id);

create table post_like
(
    id         int unsigned auto_increment
        primary key,
    post_id    int unsigned                        not null,
    user_id    int unsigned                        not null,
    created_at timestamp default CURRENT_TIMESTAMP null,
    constraint post_like_ibfk_1
        foreign key (post_id) references post (id)
            on delete cascade,
    constraint post_like_ibfk_2
        foreign key (user_id) references user (id)
            on delete cascade
);

create index post_id
    on post_like (post_id);

create index user_id
    on post_like (user_id);

create table post_section_mapping
(
    id         int unsigned auto_increment
        primary key,
    section_id int unsigned not null,
    post_id    int unsigned not null,
    constraint post_section_mapping_ibfk_1
        foreign key (section_id) references post_section (id)
            on delete cascade,
    constraint post_section_mapping_ibfk_2
        foreign key (post_id) references post (id)
            on delete cascade
);

create index post_id
    on post_section_mapping (post_id);

create index section_id
    on post_section_mapping (section_id);

create table report
(
    id          int unsigned auto_increment
        primary key,
    reported_id int unsigned                        not null,
    reporter_id int unsigned                        not null,
    reason      varchar(255)                        null,
    report_time timestamp default CURRENT_TIMESTAMP null,
    status      tinyint   default 0                 null,
    delete_at   timestamp                           null,
    type        smallint                            null,
    constraint report_ibfk_1
        foreign key (reported_id) references user (id)
            on delete cascade,
    constraint report_ibfk_2
        foreign key (reporter_id) references user (id)
            on delete cascade
);

create index reported_id
    on report (reported_id);

create index reporter_id
    on report (reporter_id);

create table social_account
(
    id       int unsigned auto_increment
        primary key,
    user_id  int unsigned not null,
    platform varchar(63)  not null,
    link     varchar(255) null,
    constraint idx_user_platform
        unique (user_id, platform),
    constraint social_account_ibfk_1
        foreign key (user_id) references user (id)
            on delete cascade
);

create index idx_email
    on user (email);

create index idx_phone
    on user (phone);

create table user_friend
(
    id             int unsigned auto_increment
        primary key,
    user_id        int unsigned       not null,
    friend_id      int unsigned       not null,
    status         tinyint  default 0 null,
    request_date   timestamp          null,
    accepted_date  timestamp          null,
    require_person int unsigned       null,
    relationship   smallint default 1 null,
    delete_at      timestamp          null,
    constraint user_friend_ibfk_1
        foreign key (user_id) references user (id)
            on delete cascade,
    constraint user_friend_ibfk_2
        foreign key (friend_id) references user (id)
            on delete cascade,
    constraint user_friend_ibfk_3
        foreign key (require_person) references user (id)
            on delete set null
);

create index friend_id
    on user_friend (friend_id);

create index require_person
    on user_friend (require_person);

create index user_id
    on user_friend (user_id);

create table user_profile
(
    user_id           int unsigned                        not null,
    display_name      varchar(50)                         null,
    avatar_id         varchar(255)                        null,
    bio               varchar(255)                        null,
    gender            tinyint                             null,
    birthday          date                                null,
    location          varchar(100)                        null,
    occupation        varchar(100)                        null,
    education         varchar(100)                        null,
    school            varchar(100)                        null,
    major             varchar(100)                        null,
    company           varchar(100)                        null,
    position          varchar(100)                        null,
    website           varchar(255)                        null,
    created_at        timestamp default CURRENT_TIMESTAMP null,
    modified_at       timestamp default CURRENT_TIMESTAMP null on update CURRENT_TIMESTAMP,
    delete_at         timestamp                           null,
    created_person    int unsigned                        null,
    modified_person   int unsigned                        null,
    modified_by_admin tinyint   default 0                 not null,
    id                int auto_increment
        primary key,
    constraint user_profile_user_id_fk
        foreign key (user_id) references user (id)
);

create table user_tag_mapping
(
    id      int unsigned auto_increment
        primary key,
    user_id int unsigned not null,
    tag_id  int unsigned not null,
    constraint user_tag_mapping_ibfk_1
        foreign key (user_id) references user (id)
            on delete cascade,
    constraint user_tag_mapping_ibfk_2
        foreign key (tag_id) references tag (id)
            on delete cascade
);

create index tag_id
    on user_tag_mapping (tag_id);

create index user_id
    on user_tag_mapping (user_id);





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

INSERT INTO post (id, author, title, content)
VALUES (1,1, 'My First Post', 'This is my first post on this platform.'),
       (2,2, 'A Day in My Life', 'Here is a glimpse into my daily routine.'),
       (3,3, 'Healthy Breakfast Ideas', 'Some healthy breakfast options.'),
       (4,4, 'How to Study Effectively', 'Tips for improving your study habits.'),
       (5,5, 'Top Movies of the Year', 'My top picks for the best movies this year.');

INSERT INTO post_history (id, author, title, content)
VALUES ( 1,1, 'My First Post (History)', 'This is the original version of my first post.'),
       ( 2,2, 'A Day in My Life (History)', 'Here is the original version of my daily routine.'),
       ( 3,3, 'Healthy Breakfast Ideas (History)', 'Some original healthy breakfast options.'),
       ( 4,4, 'How to Study Effectively (History)', 'Original tips for improving your study habits.'),
       ( 5,5, 'Top Movies of the Year (History)', 'My original top picks for the best movies this year.');

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





