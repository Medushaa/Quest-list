## Todo list CLI tool Screenshot:
![ss-v1](quest.png)

## `quest -help` 
-add string

    Add a new quest with 'title : urgency : due date' (urgency value 0 to 5)
-del int

    Type the quest index to delete (default -1)
-edit string

    Edit a quest by index & other infos like, 'index : title : urgency : dueDate'
-leftQuests

    List all the remaining quests
-list

    List all quests in a pretty table
-toggle int

    Type the quest index to toggle 'completed' (default -1)

### Examples:
    quest -add "Eat chiken : 3 : 01 Dec 2024 5PM"
    quest -edit "2:-:4:-" (changes urgency at idx 2 to 4)
    quest -list
    quest -del 1