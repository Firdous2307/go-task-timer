function startTask() {
    const taskInput = document.getElementById('task');
    const taskName = taskInput.value;

    if (!taskName) {
        alert("Please enter a task name.");
        return;
    }

    const duration = prompt("Enter the duration in seconds:");
    if (!duration || isNaN(duration)) {
        alert("Please enter a valid duration.");
        return;
    }

    fetch('/start', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/x-www-form-urlencoded',
        },
        body: new URLSearchParams({ task: taskName, duration: duration })
    })
    .then(response => response.text())
    .then(data => {
        alert(data);
        taskInput.value = ''; // Clear input
        setTimeout(() => stopTask(taskName, duration), duration * 1000);
        fetchCompletedTasks(); // Add this line
    })
    .catch(error => console.error('Error:', error));
}

function stopTask(taskName, duration) {
    fetch('/stop', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/x-www-form-urlencoded',
        },
        body: new URLSearchParams({ task: taskName, duration: duration })
    })
    .then(response => response.text())
    .then(data => {
        alert(data);
        fetchCompletedTasks();
    })
    .catch(error => console.error('Error:', error));
}


function fetchCompletedTasks() {
    fetch('/completed-tasks')
        .then(response => response.json())
        .then(tasks => {
            const taskList = document.getElementById('taskList');
            taskList.innerHTML = ''; // Clear existing list
            tasks.forEach(task => {
                const li = document.createElement('li');
                li.textContent = `${task.Name} (Duration: ${task.Duration} seconds)`;
                taskList.appendChild(li);
            });
        })
        .catch(error => {
            console.error('Error fetching completed tasks:', error);
            // Add error handling to display the error to the user
            const taskList = document.getElementById('taskList');
            taskList.innerHTML = '<li>Error fetching completed tasks. Please try again later.</li>';
        });
}

// Initial fetch when the page loads
document.addEventListener('DOMContentLoaded', fetchCompletedTasks);

// Remove the setInterval call
// setInterval(fetchCompletedTasks, 5000);