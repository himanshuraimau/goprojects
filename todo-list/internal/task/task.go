package tasks

import (
    "encoding/csv"
    "os"
    "strconv"
    "time"
    "errors"
    
    
    
)


func AddTask(description string) error {
    f, err := os.OpenFile("data/tasks.csv", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
    if err != nil {
        return err
    }
    defer f.Close()

    writer := csv.NewWriter(f)
    defer writer.Flush()

    task := Task{
        ID:          generateID(),
        Description: description,
        CreatedAt:   time.Now(),
        IsComplete:  false,
    }

    return writer.Write(task.ToCSV())
}

func (t Task) ToCSV() []string {
    return []string{
        strconv.Itoa(t.ID),
        t.Description,
        t.CreatedAt.Format(time.RFC3339),
        strconv.FormatBool(t.IsComplete),
    }
}

func generateID() int {
    tasks, err := ReadTasks()
    if err != nil {
        return 1
    }

    maxID := 0
    for _, task := range tasks {
        if task.ID > maxID {
            maxID = task.ID
        }
    }

    return maxID + 1
}

func ReadTasks() ([]Task, error) {
    f, err := os.Open("data/tasks.csv")
    if err != nil {
        return nil, err
    }
    defer f.Close()

    reader := csv.NewReader(f)
    records, err := reader.ReadAll()
    if err != nil {
        return nil, err
    }

    var tasks []Task
    for _, record := range records {
        id, err := strconv.Atoi(record[0])
        if err != nil {
            return nil, err
        }
        createdAt, err := time.Parse(time.RFC3339, record[2])
        if err != nil {
            return nil, err
        }
        isComplete, err := strconv.ParseBool(record[3])
        if err != nil {
            return nil, err
        }

        task := Task{
            ID:          id,
            Description: record[1],
            CreatedAt:   createdAt,
            IsComplete:  isComplete,
        }
        tasks = append(tasks, task)
    }

    return tasks, nil
}


func CompleteTask(taskID string) error {
    // Open the tasks CSV file for reading and writing
    file, err := os.OpenFile("data/tasks.csv", os.O_RDWR, 0666)
    if err != nil {
        return err
    }
    defer file.Close()

    reader := csv.NewReader(file)
    rows, err := reader.ReadAll()
    if err != nil {
        return err
    }

    var updatedRows [][]string
    taskFound := false

    for _, row := range rows {
        if row[0] == taskID {
            // Update the task's IsComplete status
            row[3] = "true"
            taskFound = true
        }
        updatedRows = append(updatedRows, row)
    }

    if !taskFound {
        return errors.New("task not found")
    }

    // Re-open the file for writing and save the updated rows
    file, err = os.OpenFile("data/tasks.csv", os.O_RDWR|os.O_TRUNC, 0666)
    if err != nil {
        return err
    }
    defer file.Close()

    writer := csv.NewWriter(file)
    err = writer.WriteAll(updatedRows)
    if err != nil {
        return err
    }

    return nil
}