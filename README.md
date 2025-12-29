# ACid Rss

## The Aggressively Manual RSS Reader for People Who Miss Slow Internet

Welcome to the most deliberately inconvenient RSS reader you'll ever encounter. While everyone else is building "real-time AI-curated feeds with instant push notifications and infinite scroll," we're over here making you *manually run a command* to read the news. Like some kind of caveman. With a terminal.

### What This Beautiful Disaster Does

1. **Reads RSS feeds from a text file** - We evolved. Now we use `sources.txt`. Revolutionary, I know. We went from one text file hosted on GitHub to... another text file hosted on GitHub. Progress.

2. **Generates Static HTML via Pipeline** - That's right. No dynamic content. No live updates. No WebSockets. You run the pipeline, it generates HTML files, and then you stare at them until you decide you want fresh content. Which means running the pipeline again. Manually. Because we're not animals who need automation.

3. **Publishes to GitHub Pages** - Because paying for hosting is for people with money. This costs exactly $0.00 and works perfectly fine, which is infuriating to cloud providers everywhere.

4. **Is NOT Automatic** - This is the best part. You have to *actually run the command* to fetch new articles. No cron jobs. No webhooks. No CI/CD automation running every 5 minutes burning through GitHub Actions minutes like a SpaceX rocket. When you want news, you ask for it. Like a civilized human being.

### Why Would Anyone Deliberately Make This So Inconvenient?

Excellent question. Here's the bitter truth:

**Because your internet doesn't need to be a firehose of anxiety.**

- Remember when you had to *choose* to check the news instead of having it screamed at your face 24/7?
- Remember when HackerNews wasn't vomiting 500 articles per day into your eyeballs?
- Remember when you could go outside without your phone vibrating with another "BREAKING NEWS" notification about some random startup's Series A funding?

This reader is deliberately slow. Deliberately manual. Deliberately calm. You run it when you want to read. Not when some algorithm decides you should be reading.

### Why Would Anyone Actually Use This?

Because you're one of the few people left who understands that:

- Not everything needs real-time updates (spoiler: almost nothing does)
- The FOMO economy is a scam designed to keep you scrolling
- Manually refreshing your feed is actually *liberating*
- Static files load faster than your fancy React app ever will
- You're tired of the internet treating you like a dopamine slot machine
- You want to host your RSS feeds on GitHub Pages because you're cheap and smart (there's a difference)

### Features We PROUDLY DON'T Have

- Real-time updates (refresh yourself, we're not your servants)
- Automatic fetching (make a conscious choice to read the news)
- Push notifications (your phone has enough problems)
- AI-powered feed recommendations (read what you subscribed to, not what an algorithm thinks you want)
- Social media integration (go touch grass)
- Cloud sync (it's on GitHub, git pull like an adult)
- Mobile apps (browsers work on phones, shocking I know)
- Premium subscription tier (paying money to read free RSS feeds would be peak capitalism)
- Dark mode toggle (edit the CSS, it's literally 20 lines)

### Installation

```bash
# Clone this monument to intentional friction
git clone <your-repo-url>

# Install dependencies (there's basically none because we're not psychopaths)
go mod download

# Edit sources.txt with your RSS feeds
# Yes, with a text editor. No, we don't have a GUI. 
# No, we won't build one. Vim exists for a reason.
vim sources.txt


```

### Configuration

1. Edit `sources.txt` with your RSS feed URLs (one per line, we're not monsters)
2. That's literally it. You can put a * in the end of the url to put a star class in css
3. No API keys
4. No OAuth flows
5. No environment variables
6. No .env files
7. No Docker containers
8. No Kubernetes clusters (what's wrong with you?)
9. Just a text file with URLs in it

**Pro tip:** Choose feeds that update infrequently. If you're subscribing to feeds that post 50 times a day, you've defeated the entire purpose of this project and you should reconsider your life choices.

### Performance

It's fast. Ridiculously fast. Know why? 

Because we don't do **anything**.

- No database queries (no database to query)
- No API calls (except fetching RSS feeds, which you triggered manually)
- No server-side rendering (there's no server)
- No client-side hydration (there's any JavaScript)
- No real-time updates (there's no real-time anything)
- Just files sitting on a CDN, cached by GitHub, served at the speed of light

The website loads instantly because it's literally just HTML and CSS. You know, like websites used to work before we decided everything needed to be a "web app."
