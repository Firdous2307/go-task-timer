function startTask() {
    const taskInput = document.getElementById('task');
    const taskName = taskInput.value;

    fetch('/start', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/x-www-form-urlencoded',
        },
        body: new URLSearchParams({ task: taskName })
    })
    .then(response => response.text())
    .then(data => {
        alert(data);
        taskInput.value = ''; // Clear input
    })
    .catch(error => console.error('Error:', error));
}

function stopTask() {
    const taskInput = document.getElementById('task');
    const taskName = taskInput.value;
    
    // Simulate getting duration from somewhere
    const duration = prompt("Enter the duration in seconds:");

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
        taskInput.value = ''; // Clear input
    })
    .catch(error => console.error('Error:', error));
}