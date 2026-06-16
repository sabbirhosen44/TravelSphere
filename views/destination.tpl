<div class="destination-detail-container">
    <!-- Back to Explorer Button -->
    <a href="/countries" class="back-link"><i class="fa-solid fa-arrow-left"></i> Back to Country Explorer</a>

    <div class="destination-header">
        <div class="dest-flag-box">
            <img src="{{.Country.Flag}}" alt="{{.Country.CommonName}} flag"
                style="width: 100%; height: 100%; object-fit: cover; border-radius: 4px; display: block;">
        </div>

        <div class="dest-title-box">
            <h1>{{.Country.CommonName}}</h1>
            <p class="official-name">{{.Country.Name}}</p>
            <span class="region-badge-large">{{.Country.Region}}</span>
        </div>
    </div>

    <div class="destination-grid">
        <!-- Main Details Section -->
        <div class="card dest-facts-card">
            <h2><i class="fa-solid fa-circle-info"></i> Country Facts</h2>
            <ul class="facts-list">
                <li>
                    <span class="fact-label">Capital City</span>
                    <span class="fact-value">{{.Country.Capital}}</span>
                </li>
                <li>
                    <span class="fact-label">Total Population</span>
                    <span class="fact-value">{{.Country.FormattedPopulation}}</span>
                </li>
                <li>
                    <span class="fact-label">Currencies Used</span>
                    <span class="fact-value">{{.Country.Currencies}}</span>
                </li>
                <li>
                    <span class="fact-label">Official Languages</span>
                    <span class="fact-value">{{.Country.Languages}}</span>
                </li>
                {{if .Country.LatLng}}
                <li>
                    <span class="fact-label">Geographical Coordinates</span>
                    <span class="fact-value">{{index .Country.LatLng 0 | printf "%.4f"}}&deg; N, {{index .Country.LatLng
                        1 | printf "%.4f"}}&deg; E</span>
                </li>
                {{end}}
            </ul>
        </div>

        <!-- Weather Box Section -->
        <div class="card weather-card">
            <h2><i class="fa-solid fa-cloud-sun"></i> Live Weather in {{.Country.Capital}}</h2>
            <div class="weather-brief">
                <img src="{{.Weather.ConditionIcon}}" alt="{{.Weather.Condition}}" class="weather-icon-large">
                <div class="weather-numbers">
                    <span class="weather-temp">{{.Weather.TempC}}&deg;C</span>
                    <span class="weather-text">{{.Weather.Condition}}</span>
                </div>
            </div>
            <div class="weather-details-grid">
                <div class="weather-stat-item">
                    <i class="fa-solid fa-wind"></i>
                    <div>
                        <p>Wind</p>
                        <strong>{{.Weather.WindKph}} km/h</strong>
                    </div>
                </div>
                <div class="weather-stat-item">
                    <i class="fa-solid fa-droplet"></i>
                    <div>
                        <p>Humidity</p>
                        <strong>{{.Weather.Humidity}}%</strong>
                    </div>
                </div>
            </div>
            <div class="weather-rec-box">
                <i class="fa-solid fa-compass-drafting rec-icon"></i>
                <div class="rec-text">
                    <strong>Travel Recommendation:</strong>
                    <p>{{.Weather.Recommendation}}</p>
                </div>
            </div>
        </div>
    </div>

    <div class="destination-grid second-row">
        <!-- Attractions Section -->
        <div class="card attractions-card">
            <h2><i class="fa-solid fa-monument"></i> Featured Attractions</h2>
            <div class="attractions-list-wrapper">
                {{range .Attractions}}
                <div class="attraction-detail-item">
                    <div class="attraction-dot"></div>
                    <div class="attraction-text">
                        <h4>{{.Name}}</h4>
                        <span class="kinds-label">{{.Kinds}}</span>
                    </div>
                </div>
                {{else}}
                <div class="no-attractions">No sights found in this vicinity.</div>
                {{end}}
            </div>
        </div>

        <!-- Wishlist Actions widget -->
        <div class="card wishlist-action-card">
            <h2><i class="fa-solid fa-heart"></i> Personal Travel Wishlist</h2>
            <div id="wishlist-feedback" class="wishlist-feedback-container">
                {{if .IsAddedToWishlist}}
                <div class="alert alert-info">
                    <p><i class="fa-solid fa-circle-check"></i> Already added to your wishlist!</p>
                    <a href="/wishlist" class="btn btn-outline btn-sm"><i class="fa-solid fa-heart"></i> Manage
                        Wishlist</a>
                </div>
                {{else}}
                {{if .IsAuthenticated}}
                <div class="wishlist-form-group">
                    <label for="wishlist-status">Planned Status:</label>
                    <select id="wishlist-status" class="form-control">
                        <option value="Planned">Planned (To Visit)</option>
                        <option value="Visited">Visited (Already Went)</option>
                    </select>
                </div>
                <div class="wishlist-form-group">
                    <label for="wishlist-note">Personal Travel Notes:</label>
                    <textarea id="wishlist-note" class="form-control" rows="3"
                        placeholder="Write down spots to hit, packing lists, or notes..."></textarea>
                </div>
                <button type="button" id="btn-add-wishlist" data-country="{{.Country.CommonName}}"
                    class="btn btn-primary btn-block">
                    <i class="fa-solid fa-circle-plus"></i> Add Destination to Wishlist
                </button>
                {{else}}
                <div class="unauth-wishlist-box">
                    <p class="notice-text">Interested in visiting {{.Country.CommonName}}? Authenticate to start saving
                        your trips.</p>
                    <a href="/login" class="btn btn-primary btn-block"><i class="fa-solid fa-right-to-bracket"></i>
                        Simulate Login</a>
                </div>
                {{end}}
                {{end}}
            </div>
        </div>
    </div>
</div>