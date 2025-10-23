document.addEventListener('DOMContentLoaded', () => {
    const calendarGrid = document.querySelector('.calendar-grid-custom');
    const currentMonthYearDisplay = document.getElementById('current-month-year');
    const prevMonthBtn = document.getElementById('prev-month');
    const nextMonthBtn = document.getElementById('next-month');
    const eventList = document.getElementById('event-list');

    let currentDate = new Date();

    const mockEvents = [
        { date: '2025-10-25', description: 'Project Alpha Review' },
        { date: '2025-10-28', description: 'Team Sync' },
        { date: '2025-11-01', description: 'System Upgrade' },
        { date: '2025-11-07', description: 'Client Demo Prep' },
        { date: '2025-11-15', description: 'Cyber Security Workshop' },
        { date: '2025-12-05', description: 'Annual Tech Summit' },
        { date: '2025-12-20', description: 'Holiday Party Planning' }
    ];

    function renderCalendar() {
        // Clear previous days, but preserve day names (first 7 elements)
        while (calendarGrid.children.length > 7) {
            calendarGrid.removeChild(calendarGrid.lastChild);
        }

        const year = currentDate.getFullYear();
        const month = currentDate.getMonth(); // 0-indexed

        currentMonthYearDisplay.textContent = new Date(year, month).toLocaleString('en-US', { month: 'long', year: 'numeric' });

        const firstDayOfMonth = new Date(year, month, 1);
        const lastDayOfMonth = new Date(year, month + 1, 0);
        const numDaysInMonth = lastDayOfMonth.getDate();

        // Calculate the day of the week for the first day of the month (0 for Sunday, 1 for Monday, etc.)
        const startDayIndex = firstDayOfMonth.getDay(); // Number of leading blank cells/previous month cells

        // Previous month's days
        const prevMonthLastDay = new Date(year, month, 0); // Last day of previous month
        const numDaysInPrevMonth = prevMonthLastDay.getDate();

        for (let i = startDayIndex - 1; i >= 0; i--) {
            const dayDiv = document.createElement('div');
            dayDiv.classList.add('calendar-day', 'other-month');
            dayDiv.textContent = numDaysInPrevMonth - i;
            calendarGrid.appendChild(dayDiv);
        }

        // Current month's days
        for (let i = 1; i <= numDaysInMonth; i++) {
            const dayDiv = document.createElement('div');
            dayDiv.classList.add('calendar-day', 'current-month');
            dayDiv.textContent = i;

            const fullDate = `${year}-${String(month + 1).padStart(2, '0')}-${String(i).padStart(2, '0')}`;
            if (isToday(year, month, i)) {
                dayDiv.classList.add('today');
            }

            if (hasEvent(fullDate)) {
                const eventIndicator = document.createElement('div');
                eventIndicator.classList.add('event-indicator');
                dayDiv.appendChild(eventIndicator);
            }

            calendarGrid.appendChild(dayDiv);
        }

        // Next month's days (fill to a fixed 6x7 grid = 42 cells total)
        const totalDaysRenderedSoFar = startDayIndex + numDaysInMonth;
        const totalCellsInGrid = 42; // For a consistent 6-week calendar display

        const remainingCellsToFill = totalCellsInGrid - totalDaysRenderedSoFar;

        for (let i = 1; i <= remainingCellsToFill; i++) {
            const dayDiv = document.createElement('div');
            dayDiv.classList.add('calendar-day', 'other-month');
            dayDiv.textContent = i;
            calendarGrid.appendChild(dayDiv);
        }

        renderEventList(year, month);
    }

    function isToday(year, month, day) {
        const today = new Date();
        return today.getFullYear() === year &&
               today.getMonth() === month &&
               today.getDate() === day;
    }

    function hasEvent(dateString) {
        return mockEvents.some(event => event.date === dateString);
    }

    function renderEventList(year, month) {
        eventList.innerHTML = '';
        const currentMonthEvents = mockEvents.filter(event => {
            const eventDate = new Date(event.date);
            return eventDate.getFullYear() === year && eventDate.getMonth() === month;
        }).sort((a, b) => new Date(a.date) - new Date(b.date)); // Sort by date

        if (currentMonthEvents.length === 0) {
            const li = document.createElement('li');
            li.classList.add('list-group-item', 'event-list-item'); // Add classes for styling
            li.textContent = 'No events this month.';
            eventList.appendChild(li);
        } else {
            currentMonthEvents.forEach(event => {
                const li = document.createElement('li');
                li.classList.add('list-group-item', 'event-list-item'); // Add classes for styling
                const eventDateSpan = document.createElement('span');
                eventDateSpan.classList.add('event-date');
                eventDateSpan.textContent = event.date;
                li.appendChild(eventDateSpan);
                li.appendChild(document.createTextNode(` ${event.description}`));
                eventList.appendChild(li);

                // Add click event listener to redirect to event.html
                li.addEventListener('click', () => {
                    window.location.href = 'event.html';
                });
            });
        }
    }

    prevMonthBtn.addEventListener('click', () => {
        currentDate.setMonth(currentDate.getMonth() - 1);
        renderCalendar();
    });

    nextMonthBtn.addEventListener('click', () => {
        currentDate.setMonth(currentDate.getMonth() + 1);
        renderCalendar();
    });

    renderCalendar(); // Initial render
});