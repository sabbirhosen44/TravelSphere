<div class="dashboard-container">
    <div class="dashboard-header">
        <h1>Travel Dashboard</h1>
        <p>A quick overview of your personal travel progression and saved milestones.</p>
    </div>

    <!-- Dashboard statistics panels -->
    <div id="dashboard-stats" class="stats-grid">
        <div class="stat-card total-card">
            <div class="stat-icon-wrapper">
                <i class="fa-solid fa-heart"></i>
            </div>
            <div class="stat-info">
                <h3>Total Saved</h3>
                <span class="stat-value" id="stat-total">{{.TotalCount}}</span>
                <p>Destinations in Wishlist</p>
            </div>
        </div>

        <div class="stat-card planned-card">
            <div class="stat-icon-wrapper">
                <i class="fa-solid fa-calendar-days"></i>
            </div>
            <div class="stat-info">
                <h3>Upcoming Trips</h3>
                <span class="stat-value" id="stat-planned">{{.PlannedCount}}</span>
                <p>Planned Destinations</p>
            </div>
        </div>

        <div class="stat-card visited-card">
            <div class="stat-icon-wrapper">
                <i class="fa-solid fa-map-location-dot"></i>
            </div>
            <div class="stat-info">
                <h3>Milestones</h3>
                <span class="stat-value" id="stat-visited">{{.VisitedCount}}</span>
                <p>Visited Destinations</p>
            </div>
        </div>
    </div>

    <!-- Quick breakdown layout lists -->
    <div class="dashboard-grid">
        <div class="card panel-card">
            <h2><i class="fa-solid fa-plane-departure"></i> Next Steps (Upcoming)</h2>
            <div class="dashboard-list" id="dashboard-planned-list">
                <p class="helper-text">Add items or update statuses on your Wishlist page to populate this panel.</p>
            </div>
        </div>

        <div class="card panel-card">
            <h2><i class="fa-solid fa-award"></i> Checked Off (Visited)</h2>
            <div class="dashboard-list" id="dashboard-visited-list">
                <p class="helper-text">Your achievements will appear here after marking destinations as visited.</p>
            </div>
        </div>
    </div>
</div>