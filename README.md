# windows update disabler

### Build the source by running
```bash
go build -ldflags -H=windowsgui -o wud.exe wud.go
```

### Usage
#### Running a Program with Administrator Privileges on Startup

- This guide explains how to use Task Scheduler in Windows to run a program with administrator privileges at startup without triggering a User Account Control (UAC) prompt every time.

## Steps

### 1. Open Task Scheduler
- Press `Win + S` and search for **Task Scheduler**, then open it.

### 2. Create a New Task
- In the Task Scheduler window, click on **Create Task** in the right-hand pane.

### 3. General Tab
- Name the task (e.g., **wud**).
- Check **Run with highest privileges** to ensure it runs as administrator.
- Choose **Configure for:** Windows 10 (or your Windows version).

### 4. Triggers Tab
- Click **New...**.
- Set **Begin the task** to **At log on**.
- Leave the settings for "Any user" or configure it for a specific user.
- Click **OK**.

### 5. Actions Tab
- Click **New...**.
- Set **Action** to **Start a program**.
- Click **Browse** and select the `wud.exe` file of your program.
- Click **OK**.

### 6. Conditions Tab
- Uncheck **Start the task only if the computer is on AC power** (if this option is checked).

### 7. Finish the Task
- Click **OK** to save the task.

### 8. Test the Task
- Restart your computer to verify that the program runs automatically with admin rights.

## Conclusion
Following these steps will ensure that your program runs with the necessary privileges at startup without prompting for UAC each time.