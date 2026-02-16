# **Gator**

Gator is a command-line RSS feed aggregator built in **Go**. It allows you to manage multiple users, subscribe to RSS feeds, and aggregate posts into a local PostgreSQL database.

---

### **FEATURES**
* **User Management**: Register and switch between different user profiles.
* **Feed Subscription**: Add new RSS feeds and follow/unfollow existing ones.
* **Background Aggregation**: Fetch and store posts from your followed feeds.
* **Post Browsing**: View the latest posts directly in your terminal.

---

### **PREREQUISITES**
* **Go**: Version 1.21 or higher.
* **PostgreSQL**: A running instance to store data.
* **Goose**: A database migration tool.

---

### **INSTALLATION**
1. **Clone the repository**:
```bash
git clone [https://github.com/Struhola/Gator.git](https://github.com/Struhola/Gator.git)
cd Gator
```
2. **Install the binary**:
```bash
go install github.com/Struhola/Gator@latest
```

Ensure your $GOPATH/bin is in your system's PATH to run the gator command globally.

### **SETUP**
1. **Database Configuration**

    Create a new PostgreSQL database:
```SQL
CREATE DATABASE gator;
```

2. **Configuration File**

    Gator looks for a configuration file located at ~/.gatorconfig.json. Create this file and add your database connection string:
```JSON
{
  "db_url": "postgres://username:password@localhost:5432/gator",
  "current_user_name": ""
}
```
Replace username and password with your actual PostgreSQL credentials.

3. **Database Migrations**
    Run the migrations to set up the necessary table schema:
```Bash
goose -dir sql/schema postgres "postgres://username:password@localhost:5432/gator" up
```


### **USAGE**

**User Commands**
* **Register a user**: gator register <name>
* **Login**: gator login <name>
* **List users**: gator users

**Feed Management:**
* **Add a feed**: gator addfeed <name> <url>
* **List all feeds**: gator feeds
* **Follow a feed**: gator follow <url>
* **Unfollow a feed**: gator unfollow <url>
* **List followed feeds**: gator following

**Aggregation & Browsing:**
* **Start Aggregator**: gator agg <time_interval>
    Example: gator agg 1m (checks for new posts every minute).
* **Browse Posts**: gator browse <optional_limit>
    Example: gator browse 5 (shows the 5 most recent posts).

### **Development:**
To run the project in development mode without installing:
```Bash
go run . <command> <args>
```
To reset the database (delete all data):
```Bash
gator reset
```