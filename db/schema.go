package db

import "time"

const enable_foreign_keys string = `PRAGMA foreign_keys = ON`

const user_table string = `
CREATE TABLE IF NOT EXISTS user (
    id INTEGER PRIMARY KEY,
    username TEXT,
    displayName TEXT,
    apiKey TEXT
    )`

const stream_table string = `
CREATE TABLE IF NOT EXISTS stream (
    id INTEGER PRIMARY KEY,
    twid INTEGER,
    title TEXT,
    startedAt DATETIME,
    endedAt DATETIME,
    userId INTEGER NOT NULL,
    first_userId INTEGER,
    qotdId INTEGER,
    FOREIGN KEY (userId)
    REFERENCES user (id)
        ON UPDATE CASCADE
        ON DELETE CASCADE
    FOREIGN KEY (first_userId)
    REFERENCES user (id)
        ON UPDATE SET NULL
        ON DELETE SET NULL
    FOREIGN KEY (qotdId)
    REFERENCES question (id)
        ON UPDATE SET NULL
        ON DELETE SET NULL
    )`

const question_table string = `
CREATE TABLE IF NOT EXISTS question (
    id INTEGER PRIMARY KEY,
    text TEXT UNIQUE
    )`

type User struct {
	ID          int
	Username    string
	DisplayName string
	APIKey      *string
}

type Stream struct {
	ID          int
	TWID        *int
	UserId      int
	Title       *string
	StartedAt   time.Time
	EndedAt     *time.Time
	QOTDId      *int
	FirstUserId *int
}

type Question struct {
	ID        int64  `json:"id"`
	Text      string `json:"text"`
	Disabled  bool   `json:"disabled"`
	SkipCount int    `json:"skipCount"`
}

const userSeed = `
INSERT INTO user (id, username, displayName, apiKey)
VALUES
(31568083,	"soulxburn", "SouLxBurN", "abcd1234"),
(236797464,	"caleb_dev", "caleb_dev", null)
`
const questionSeed = `
INSERT INTO question(text)
VALUES
("What's the best piece of advice you've ever received?"),
("If you could have any superpower, what would it be and why?"),
("What's your favorite way to relax after a long day?"),
("If you could travel anywhere in the world right now, where would you go?"),
("What's the last book you read or movie you watched that you couldn't put down?"),
("What's the most memorable meal you've ever had?"),
("If you could have dinner with any historical figure, who would it be and why?"),
("What's the most adventurous thing you've ever done?"),
("If you could instantly learn any skill, what would it be?"),
("What's your favorite way to spend a weekend?"),
("What's the most interesting thing you've learned recently?"),
("If you could only eat one type of cuisine for the rest of your life, what would it be?"),
("What's your favorite hobby or pastime?"),
("If you could meet any fictional character, who would it be and why?"),
("What's the best concert or live performance you've ever attended?"),
("What's the most challenging thing you've ever accomplished?"),
("If you could have a conversation with your future self, what would you ask?"),
("What's one thing you've always wanted to try but haven't had the opportunity yet?"),
("If you could choose a new name for yourself, what would it be?"),
("What's your favorite quote or mantra that inspires you?"),
("If you could magically become fluent in any language, which one would you choose?"),
("What's the best piece of advice you would give to your younger self?"),
("What's the most beautiful place you've ever visited?"),
("If you were stranded on a desert island, what three items would you want to have with you?"),
("What's the most interesting job you've ever had?"),
("If you could be an expert in any field, what would it be?"),
("What's the most courageous thing you've ever done?"),
("If you could have dinner with any celebrity, who would you choose?"),
("What's your favorite way to stay active and fit?"),
("If you could invent a new technology, what problem would it solve?"),
("What's your favorite board game or card game?"),
("If you could time travel, which era or event would you visit?"),
("What's one skill or hobby you've always wanted to learn but haven't yet?"),
("If you could have any animal as a pet, what would you choose?"),
("What's your favorite season and why?"),
("If you could trade lives with someone for a day, who would it be?"),
("What's the most memorable gift you've ever received?"),
("If you could eliminate one chore or task from your life forever, what would it be?"),
("What's your favorite way to show kindness to others?"),
("If you could witness any historical event, what would it be?"),
("What's your go-to karaoke song?"),
("If you could have a dinner party with three famous people, living or deceased, who would you invite?"),
("What's your favorite way to celebrate your birthday?"),
("If you could master any musical instrument, which one would you choose?"),
("What's your favorite way to de-stress or unwind?"),
("If you could be a character in a book or movie, who would you be and why?"),
("What's the most valuable lesson you've learned from a failure or mistake?"),
("If you could live in any fictional world, where would you choose to live?"),
("What's the best piece of advice you would give to someone starting a new job or career?"),
("If you could have a conversation with any historical figure, who would it be and what would you ask them?"),
("What's your favorite outdoor activity or sport?"),
("If you could have an unlimited supply of one thing, what would it be?"),
("What's your favorite way to practice self-care?"),
("If you could be an expert in any form of art, what would you choose?"),
("What's the most interesting fact you've learned recently?"),
("If you could have a personal chef, what type of food would you want them to prepare?"),
("What's the most memorable concert or live performance you've ever attended?"),
("If you could spend a day with any fictional character, who would it be and what would you do together?"),
("What's your favorite way to give back to your community?"),
("If you could have a conversation with any animal, which one would you choose and what would you ask?"),
("What's your favorite type of dessert?"),
("If you could have any career, regardless of qualifications or training, what would you choose?"),
("What's the best piece of advice you've ever received about relationships?"),
("If you could have a home anywhere in the world, where would it be?"),
("What's the most interesting documentary you've ever watched?"),
("If you could bring back any fashion trend, what would it be?"),
("What's your favorite way to start your day?"),
("If you could meet any famous athlete, who would you choose?"),
("What's the most thrilling adventure or activity you've ever experienced?"),
("If you could have any vehicle, whether it's practical or not, what would you choose?"),
("What's your favorite way to express your creativity?"),
("If you could solve one global problem, what would it be and why?"),
("What's your favorite way to spend quality time with friends or family?"),
("If you could have any fictional character as a best friend, who would it be and why?"),
("What's the most interesting historical fact you know?"),
("If you could have a conversation with any person, dead or alive, who would it be and why?"),
("What's your favorite way to learn something new?"),
("If you could have a lifetime supply of any snack, what would you choose?"),
("What's the most unusual food you've ever tried?"),
("If you could be a contestant on any game show, which one would you pick?"),
("What's your favorite way to stay motivated and productive?"),
("If you could visit any planet in the solar system, which one would you choose and why?"),
("What's the most inspiring book you've ever read?"),
("If you could witness any natural phenomenon, what would it be?"),
("What's your favorite type of music or favorite band/artist?"),
("If you could bring any fictional character to life, who would it be and why?"),
("What's the best piece of advice you would give to someone about pursuing their dreams?"),
("If you could have dinner with any family member, living or deceased, who would you choose and why?"),
("What's your favorite way to explore new places or cities?"),
("If you could have any animal talent or ability, what would it be?"),
("What's the most interesting historical landmark you've ever visited?"),
("If you could learn a new language instantly, which one would you choose and why?"),
("What's your favorite way to overcome a challenge or obstacle?"),
("If you could have any job for a day, just to try it out, what would it be?"),
("What's the most memorable concert or live performance you've ever been a part of (as a performer or audience)?"),
("If you could be an expert in any form of dance, which one would you choose?"),
("What's your favorite way to connect with nature?"),
("If you could have any famous artist create a portrait of you, who would you choose?"),
("What's the best piece of advice you would give to someone about maintaining a healthy lifestyle?"),
("If you could be a character in a video game, who would you be and why?"),
("What's your favorite way to give yourself a mental or emotional boost?"),
("If you could have any fictional technology or gadget from a movie, what would it be?"),
("What's the most interesting animal fact you know?"),
("If you could have a conversation with any musician, living or deceased, who would you choose and why?"),
("What's your favorite way to stay organized and manage your time effectively?"),
("If you could visit any historical era as an observer, which one would you choose and why?"),
("What's the most memorable sporting event you've ever attended?"),
("If you could have any mythical creature as a pet, what would you choose?"),
("What's your favorite way to express gratitude?"),
("If you could have a talent in any form of performing arts, which one would you choose?"),
("What's the most interesting scientific discovery or concept you've come across?"),
("If you could have any job in the world, regardless of qualifications or salary, what would you choose?"),
("What's your favorite way to enjoy a rainy day?"),
("If you could have any historical artifact in your possession, what would it be?"),
("What's the best piece of advice you would give to someone about pursuing their passions?"),
("If you could have any animal as a sidekick, which one would you choose?"),
("What's your favorite way to unwind and relax on a Friday night?"),
("If you could learn instantly and master any instrument, which one would you choose?"),
("What's the most interesting historical event you've ever studied?"),
("If you could have a dinner party with any three people, living or deceased, who would you invite?"),
("If you could live in any fictional universe, which one would you choose?"),
("What's the most unique or unusual talent you possess?"),
("What's your favorite way to stay motivated and inspired?"),
("If you could have any job in the world, without worrying about money or qualifications, what would you choose?"),
("If you could have a conversation with any historical figure, who would you choose and why?"),
("What's your favorite way to spend a lazy Sunday morning?"),
("If you could instantly learn any language, which one would you choose and why?"),
("If you could have a superpower for a day, what would you choose and how would you use it?"),
("What's your favorite way to start your day on a positive note?"),
("If you could have any famous artist create a portrait of you, who would you choose and why?"),
("If you could bring back any fashion trend from the past, which one would it be?"),
("If you could travel back in time and give your younger self advice, what would it be?"),
("If you could have any animal's ability or trait, which one would you choose?"),
("What's your favorite way to practice self-care and prioritize your well-being?"),
("If you could have a conversation with your future self, what advice or questions would you share?"),
("What's your favorite way to explore and discover new things in your city or town?"),
("If you could have dinner with any fictional character, who would you choose and why?"),
("If you could have any career for a day, just to try it out, what would it be?"),
("What's your favorite way to celebrate special occasions?"),
("If you could bring one extinct animal back to life, which one would you choose?"),
("What's the most interesting piece of trivia about yourself?"),
("If you could learn any form of dance instantly, which one would you choose?"),
("What's your favorite way to spark your creativity or find inspiration?"),
("What's the most interesting historical landmark or site you've ever visited?"),
("If you could have any technological innovation become a reality, what would you choose?"),
("What's your favorite way to spend quality time with your loved ones?"),
("If you could have any famous athlete's skills for a day, who would you choose?"),
("What's your favorite way to connect with nature and the outdoors?"),
("If you could be an expert in any form of art, which one would you choose?"),
("What's the most interesting piece of information or news you've come across recently?"),
("What's the most interesting animal fact or behavior you know?"),
("If you could have a personal chef for a week, what type of food would you want them to prepare?"),
("What's your favorite way to stay organized and manage your tasks effectively?"),
("If you could have any historical artifact in your possession, what would it be and why?"),
("What's your favorite way to enjoy a sunny day?"),
("If you could be a contestant on any reality TV show, which one would you pick?"),
("What's the most memorable piece of advice you've ever received about success and achieving goals?"),
("If you could have any fictional character's wardrobe, whose would you choose?"),
("What's your favorite way to give yourself a mental break and relax your mind?"),
("If you could have any celebrity as your mentor, who would you choose and why?"),
("What's the most interesting fact or story about your family history?"),
("If you could have a conversation with any artist, living or deceased, who would you choose and why?"),
("What's your favorite way to challenge yourself and step out of your comfort zone?"),
("If you could have any magical ability or power, what would it be?"),
("What's the most interesting scientific discovery or concept you've come across recently?"),
("If you could visit any country in the world, where would you go and what would you do?"),
("What's your favorite way to give yourself a creative outlet?"),
("If you could have a conversation with any character from a TV show, who would you choose and why?"),
("What's the most memorable event or celebration you've ever attended?"),
("If you could have any famous artist's artwork displayed in your home, whose would you choose?"),
("What's your favorite way to overcome a challenge or obstacle in your life?"),
("If you could be a master of any form of martial arts, which one would you choose?"),
("What's the most interesting piece of technology you've come across recently?"),
("If you could have any profession for a day, just to experience it, what would it be?"),
("What's your favorite way to relax and recharge during a vacation?"),
("What's the most interesting piece of trivia you know about a famous landmark or monument?"),
("If you could have a conversation with any mythical creature, which one would you choose and what would you ask?"),
("What's your favorite way to learn new things or acquire new skills?"),
("If you could have any famous actor or actress star in a movie about your life, who would you choose?"),
("What's the most adventurous or daring thing you've ever done while traveling?"),
("If you could have a conversation with any historical leader or ruler, who would you choose and why?"),
("What's your favorite way to express your gratitude towards others?"),
("If you could have any celebrity's fashion sense, whose wardrobe would you choose?"),
("What's the most interesting natural phenomenon you've ever witnessed?"),
("If you could have a conversation with any musician from a different era, who would you choose and why?"),
("What's your favorite way to overcome creative blocks or writer's block?"),
("If you could attend any major sporting event, which one would you choose and why?"),
("What's the most memorable piece of advice you've received from a family member or loved one?"),
("If you could have any famous writer or author write your biography, who would you choose and why?"),
("What's your favorite way to learn about different cultures and traditions?"),
("If you could be a character in a fairy tale, who would you be and why?"),
("What's the most interesting historical fact or event from your own country?"),
("If you could have a conversation with any comedian, living or deceased, who would you choose and why?"),
("What's your favorite way to stay mentally sharp and challenge your mind?"),
("If you could have any fictional vehicle or transportation device, what would it be?"),
("What's the most memorable adventure or expedition you've ever been on?"),
("If you could have any job related to the arts, which one would you choose?"),
("What's your favorite way to express your creativity in everyday life?"),
("If you could have a conversation with any philosopher, who would you choose and why?"),
("What's the most interesting piece of trivia you know about a famous celebrity?"),
("If you could have any superpower related to knowledge or learning, what would it be?"),
("What's your favorite way to celebrate personal achievements or milestones?"),
("If you could be a character in a famous painting, which one would you choose?"),
("What's the most memorable concert or live performance you've ever watched online?"),
("If you could have any historical document or manuscript in your possession, what would it be?"),
("What's your favorite way to enjoy nature and the great outdoors?"),
("If you could have any famous chef cook a meal for you, who would you choose and why?"),
("What's the most interesting fact you've learned about space or the universe?"),
("If you could have a conversation with any leader or influential figure, who would you choose and why?"),
("What's your favorite way to boost your confidence and self-esteem?"),
("If you could have any famous athlete's skills in a specific sport, which sport and athlete would you choose?"),
("What's the most memorable gift you've ever given to someone?"),
("If you could solve one social issue or challenge, what would it be and why?"),
("What's your favorite way to connect with others and build meaningful relationships?"),
("If you could have any historical artifact from ancient civilizations, what would you choose and why?"),
("What's your favorite way to enjoy a sunset or sunrise?"),
("If you could be a contestant on any cooking show, which one would you pick?"),
("What's the most valuable piece of advice you would give to your future self?"),
("If you could bring any fictional creature to life, which one would you choose and why?"),
("What's your favorite way to learn about history and explore the past?"),
("If you could have a conversation with any business tycoon or entrepreneur, who would you choose and why?"),
("What's the most interesting cultural tradition or festival you've ever experienced?"),
("If you could have any famous actor or actress portray you in a movie, who would you choose and why?"),
("What's your favorite way to escape reality and immerse yourself in a different world (books, movies, video games, etc.)?"),
("If you could have any technology invented to simplify your daily life, what would it be?"),
("What's the most interesting piece of trivia you know about your favorite hobby or interest?"),
("If you could have a conversation with any influential scientist, who would you choose and why?"),
("What's your favorite way to volunteer and make a positive impact in your community?"),
("If you could be a character in a historical novel, which time period would you choose?"),
("What's the most memorable event or party you've ever organized?"),
("If you could have any famous artist create a mural on your home's exterior, who would you choose?"),
("What's your favorite way to stay positive and maintain a optimistic mindset?"),
("If you could be an expert in any form of visual arts, which one would you choose?"),
("What's the most interesting scientific experiment or study you've ever heard of?"),
("If you could have any famous singer perform at your wedding, who would you choose?"),
("What's your favorite way to practice mindfulness and be present in the moment?"),
("If you could have any fictional vehicle from a sci-fi movie, which one would you choose?"),
("What's the most interesting animal behavior or adaptation you've learned about?"),
("If you could have a conversation with any historical inventor or scientist, who would you choose and why?"),
("What's your favorite way to surprise or delight someone special in your life?"),
("If you could have any profession from a different era, which one would you choose?"),
("What's the most memorable adventure or excursion you've ever been on with friends or family?"),
("If you could have any superhero's costume, which one would you choose?"),
("What's your favorite way to stimulate your mind and engage in intellectual discussions?"),
("If you could have a conversation with any musician from a different genre, who would you choose and why?"),
("What's the most interesting piece of trivia you know about a historical figure?"),
("If you could have any famous actor or actress as your mentor, who would you choose and why?"),
("What's your favorite way to give back to the environment and practice sustainability?"),
("If you could attend any major award show, which one would you choose and why?"),
("What's the most interesting fact or piece of trivia you know about a famous landmark?"),
("If you could have any magical item from a fantasy book or movie, what would it be?"),
("What's your favorite way to stimulate your creativity and generate new ideas?"),
("If you could have a conversation with any literary character, who would you choose and why?"),
("What's the most memorable event or celebration you've ever hosted?"),
("If you could have any famous artist create a sculpture of you, who would you choose and why?"),
("What's your favorite way to relax and find inner peace?"),
("If you could have any character from a TV show as your best friend, who would you choose and why?"),
("What's the most interesting piece of trivia you know about a famous historical battle?"),
("If you could have a conversation with any philosopher or thinker, living or deceased, who would you choose and why?"),
("What's your favorite way to challenge yourself physically and stay fit?"),
("If you could have any fictional power from a comic book, what would it be?"),
("What's your favorite way to support and uplift others in your community?"),
("If you could attend any major music festival, which one would you choose and why?"),
("What's the most interesting fact or piece of trivia you know about a famous work of art?"),
("If you could have a conversation with any historical figure from your country, who would you choose and why?"),
("What's your favorite way to stay motivated and achieve your goals?"),
("If you could have any famous actor or actress as your co-star in a movie, who would you choose and why?"),
("What's the most memorable event or celebration you've ever attended as a guest?"),
("If you could have any famous architect design your dream home, who would you choose and why?"),
("What's your favorite way to explore and learn about different cuisines and culinary traditions?"),
("If you could have a conversation with any sports legend, who would you choose and why?"),
("What's the most interesting fact or piece of trivia you know about a historical invention?"),
("If you could have any musical instrument masterfully played for you, which one would you choose?"),
("What's your favorite way to practice gratitude and appreciate the little things in life?"),
("If you could attend any major film premiere, which movie would you choose and why?"),
("What's the most interesting historical fact or event from a country you've always wanted to visit?"),
("If you could have a conversation with any influential figure from the fashion industry, who would you choose and why?"),
("What's your favorite way to support local businesses and artisans?"),
("If you could have any famous actor or actress narrate your life story, who would you choose and why?"),
("What's the most memorable event or celebration you've ever organized as a host?"),
("If you could have any famous fashion designer create a custom outfit for you, who would you choose and why?"),
("What's your favorite way to embrace and appreciate different cultures and diversity?"),
("If you could have a conversation with any scientist or researcher, who would you choose and why?"),
("What's the most interesting fact or piece of trivia you know about a famous historical figure?"),
("If you could have any famous athlete as your personal coach, who would you choose and why?"),
("What's your favorite way to immerse yourself in a different time period or era?"),
("If you could have any famous actor or actress perform a monologue written specifically for you, who would you choose and why?"),
("What's the most memorable event or celebration you've attended as a participant or performer?"),
("If you could have any famous artist create a customized tattoo for you, who would you choose and why?"),
("What's your favorite way to give back and support charitable causes?"),
("If you could attend any major fashion show, which one would you choose and why?"),
("What's the most interesting fact or piece of trivia you know about a famous historical document?"),
("If you could have a conversation with any influential figure from the technology industry, who would you choose and why?"),
("What's your favorite way to celebrate and appreciate the achievements of others?"),
("If you could have any famous actor or actress as your scene partner in a play, who would you choose and why?"),
("What's the most memorable event or celebration you've ever participated in as an organizer or planner?"),
("If you could have any famous architect design a landmark for your city, who would you choose and why?"),
("What's your favorite way to explore and experience different culinary traditions from around the world?"),
("If you could have a conversation with any legendary sports coach, who would you choose and why?"),
("What's the most interesting fact or piece of trivia you know about a famous historical discovery?"),
("If you could have any musical instrument played by a world-renowned musician, which one would you choose?"),
("What's your favorite way to express appreciation and gratitude towards others?"),
("If you could attend any major theater production, which one would you choose and why?"),
("What's the most interesting historical fact or event from a country you've always been fascinated by?"),
("If you could have a conversation with any influential figure from the beauty industry, who would you choose and why?"),
("What's your favorite way to support local artists and creatives in your community?"),
("If you could have any famous actor or actress perform a dramatic reading of your favorite book, who would you choose and why?"),
("What's the most memorable event or celebration you've ever attended as a guest of honor?"),
("If you could have any famous fashion designer create a one-of-a-kind outfit for you, who would you choose and why?"),
("What's your favorite way to embrace cultural diversity and learn about different customs and traditions?"),
("If you could have a conversation with any renowned scientist, who would you choose and why?"),
("What's the most interesting fact or piece of trivia you know about a famous historical period?"),
("If you could have any famous athlete as your training partner, who would you choose and why?"),
("What's your favorite way to immerse yourself in a different cultural experience or festival?"),
("If you could have any famous actor or actress star in a play written by you, who would you choose and why?"),
("What's the most memorable event or celebration you've ever organized as a host with a specific theme?"),
("If you could have any famous architect design your dream workplace or office, who would you choose and why?"),
("What's your favorite way to explore different flavors and cuisines through food tastings or culinary tours?"),
("If you could have a conversation with any influential figure from the automotive industry, who would you choose and why?"),
("What's the most interesting fact or piece of trivia you know about a famous historical artifact?"),
("If you could have any musical instrument played by a world-class orchestra, which one would you choose?"),
("What's your favorite way to express gratitude and acknowledge the contributions of others?"),
("If you could attend any major art exhibition, which one would you choose and why?"),
("What's the most interesting historical fact or event from a country you've always wanted to explore further?"),
("If you could have a conversation with any renowned inventor or engineer, who would you choose and why?"),
("What's your favorite way to support local sports teams and athletes in your community?"),
("If you could have any famous actor or actress perform a monologue written by you, who would you choose and why?"),
("What's the most memorable event or celebration you've ever attended related to a specific hobby or interest?"),
("If you could have any famous architect design a park or public space in your city, who would you choose and why?"),
("What's your favorite way to try new recipes and experiment with different cooking techniques?"),
("If you could have a conversation with any influential figure from the aviation industry, who would you choose and why?"),
("What's the most interesting fact or piece of trivia you know about a famous historical figure's personal life?"),
("If you could have any musical instrument played by a world-renowned band or ensemble, which one would you choose?"),
("What's your favorite way to express appreciation and gratitude towards nature and the environment?"),
("If you could attend any major dance performance or show, which one wouldyou choose and why?"),
("What's the most interesting historical fact or event from a country you've always been curious about?"),
("If you could have a conversation with any influential figure from the film industry, who would you choose and why?"),
("What's your favorite way to support local theater productions and performing arts organizations?"),
("If you could have any famous actor or actress star in a movie that you write and direct, who would you choose and why?"),
("What's the most memorable event or celebration you've ever attended as a guest of honor, related to a specific interest or passion of yours?"),
("If you could have any famous architect design a museum or gallery dedicated to your favorite subject, who would you choose and why?"),
("What's your favorite way to explore different culinary traditions and cuisines through food festivals or international food fairs?"),
("If you could have a conversation with any influential figure from the gaming industry, who would you choose and why?"),
("What's the most interesting fact or piece of trivia you know about a famous historical figure's achievements or contributions?"),
("If you could have any musical instrument played by a renowned musician in a private concert, which one would you choose?")
`
