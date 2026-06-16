<div class="home-hero">
    <div class="hero-content">
        <h1>Discover Your Next Adventure</h1>
        <p>Explore countries, check local live weather, view popular sights, and construct your custom trip wishlist.
        </p>

        <div class="search-box-wrapper">
            <div class="search-input-group">
                <i class="fa-solid fa-magnifying-glass search-icon"></i>
                <input type="text" id="home-search"
                    placeholder="Search destinations (e.g. France, Japan, Bangladesh...)" autocomplete="off">
                <div id="home-search-spinner" class="spinner hidden"><i class="fa-solid fa-spinner fa-spin"></i></div>
            </div>

            <div id="home-search-results" class="ajax-search-dropdown hidden"></div>
        </div>
    </div>
</div>

<div class="section-container featured-section">
    <div class="section-header">
        <h2>Featured Countries</h2>
        <a href="/countries" class="section-link">View All <i class="fa-solid fa-arrow-right"></i></a>
    </div>

    <div class="countries-grid">
        {{range .FeaturedCountries}}
        <div class="country-card">
            <div class="country-card-header">
                <img src="{{.Flag}}" alt="{{.CommonName}} flag" class="country-flag-large"
                    style="height: 2em; width: auto; object-fit: cover; border-radius: 4px;">
            </div>
            <div class="country-card-body">
                <h3>{{.CommonName}}</h3>
                <span class="country-region-badge">{{.Region}}</span>
                <p><strong>Capital:</strong> {{.Capital}}</p>
                <p><strong>Population:</strong> {{.FormattedPopulation}}</p>
            </div>
            <div class="country-card-footer">
                <a href="/countries/{{.Slug}}" class="btn btn-primary btn-block">Explore Destination</a>
            </div>
        </div>
        {{end}}
    </div>
</div>

<div class="section-container attractions-section">
    <div class="section-header">
        <h2>Popular Attractions </h2>
    </div>

    <div class="attractions-grid">
        {{range .PopularAttractions}}
        <div class="attraction-item">
            <div class="attraction-icon-box">
                <i class="fa-solid fa-monument"></i>
            </div>
            <div class="attraction-details">
                <h3>{{.Name}}</h3>
                <span class="kinds-badge">{{.Kinds}}</span>
                <p>Curated tourist spot for international sightseers.</p>
            </div>
        </div>
        {{end}}
    </div>
</div>