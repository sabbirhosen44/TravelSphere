<header class="main-header">
    <div class="header-container">
        <a href="/" class="logo">
            <i class="fa-solid fa-earth-americas logo-icon"></i>
            <span>TravelSphere</span>
        </a>

        <nav class="nav-links">
            <a href="/" class="{{if eq .ActivePage " home"}}active{{end}}">
                <i class="fa-solid fa-house"></i> Home
            </a>
            <a href="/countries" class="{{if eq .ActivePage " countries"}}active{{end}}">
                <i class="fa-solid fa-compass"></i> Explorer
            </a>
            <a href="/wishlist" class="{{if eq .ActivePage " wishlist"}}active{{end}}">
                <i class="fa-solid fa-heart"></i> Wishlist
            </a>
            <a href="/dashboard" class="{{if eq .ActivePage " dashboard"}}active{{end}}">
                <i class="fa-solid fa-chart-line"></i> Dashboard
            </a>
        </nav>

        <div class="auth-actions">
            {{if .IsAuthenticated}}
            <span class="user-badge">
                <i class="fa-solid fa-user-astronaut"></i> Simulated Traveler
            </span>
            <a href="/logout" class="btn btn-outline btn-sm logout-btn">
                <i class="fa-solid fa-right-from-bracket"></i> Logout
            </a>
            {{else}}
            <a href="/login" class="btn btn-primary btn-sm login-btn">
                <i class="fa-solid fa-right-to-bracket"></i> Simulate Login
            </a>
            {{end}}
        </div>
    </div>
</header>