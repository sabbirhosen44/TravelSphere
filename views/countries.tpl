<div class="explorer-hero">
    <div class="explorer-hero-content">
        <h1>Country Explorer</h1>
        <p>Browse nations, filter by geographical region, and search for your dream travel destinations.</p>
    </div>
</div>

<div class="explorer-container">
    <div class="filters-bar">
        <div class="search-input-group filter-search">
            <i class="fa-solid fa-magnifying-glass search-icon"></i>
            <input type="text" id="explorer-search" placeholder="Search by name, capital, or code...">
        </div>

        <div class="region-select-group">
            <label for="explorer-region"><i class="fa-solid fa-filter"></i> Region:</label>
            <select id="explorer-region">
                <option value="">All Regions</option>
                <option value="Europe">Europe</option>
                <option value="Asia">Asia</option>
                <option value="Americas">Americas</option>
                <option value="Africa">Africa</option>
                <option value="Oceania">Oceania</option>
            </select>
        </div>

        <div id="explorer-spinner" class="spinner hidden">
            <i class="fa-solid fa-spinner fa-spin"></i> Loading...
        </div>
    </div>

    <div id="country-results" class="countries-grid">
        {{if .Error}}
        <div class="alert alert-danger">{{.Error}}</div>
        {{else}}
        {{range .Countries}}
        <div class="country-card">
            <div class="country-card-header">
                <img src="{{.Flag}}" alt="{{.CommonName}} flag" class="country-flag-large"
                    style="height: 40px; width: auto; object-fit: cover; border-radius: 4px;">
                <span class="country-region-badge">{{.Region}}</span>
            </div>
            <div class="country-card-body">
                <h3>{{.CommonName}}</h3>
                <p><strong>Capital:</strong> {{.Capital}}</p>
                <p><strong>Population:</strong> {{.FormattedPopulation}}</p>
                <p><strong>Currency:</strong> {{.Currencies}}</p>
                <p><strong>Languages:</strong> {{.Languages}}</p>
            </div>
            <div class="country-card-footer">
                <a href="/countries/{{.Slug}}" class="btn btn-primary btn-block">Explore Destination</a>
            </div>
        </div>
        {{else}}
        <div class="no-results">
            <i class="fa-solid fa-earth-europe"></i>
            <p>No countries matching filters were found.</p>
        </div>
        {{end}}
        {{end}}
    </div>

    <div id="pagination-container" class="pagination-wrapper">
        {{if .TotalPages}}
        {{if gt .TotalPages 1}}
        <nav class="pagination-nav">
            {{if .HasPrev}}
            <a href="/countries?page={{.PrevPage}}{{if .CurrentSearch}}&search={{.CurrentSearch}}{{end}}{{if .CurrentRegion}}&region={{.CurrentRegion}}{{end}}"
                class="btn-pagination prev-page" data-page="{{.PrevPage}}"><i class="fa-solid fa-chevron-left"></i>
                Previous</a>
            {{end}}

            <span class="page-info">Page <strong id="current-page-num">{{.CurrentPage}}</strong> of <strong
                    id="total-pages-num">{{.TotalPages}}</strong></span>

            {{if .HasNext}}
            <a href="/countries?page={{.NextPage}}{{if .CurrentSearch}}&search={{.CurrentSearch}}{{end}}{{if .CurrentRegion}}&region={{.CurrentRegion}}{{end}}"
                class="btn-pagination next-page" data-page="{{.NextPage}}">Next <i
                    class="fa-solid fa-chevron-right"></i></a>
            {{end}}
        </nav>
        {{end}}
        {{end}}
    </div>
</div>