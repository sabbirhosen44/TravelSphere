<div class="wishlist-container">
    <div class="wishlist-header">
        <h1>My Travel Wishlist</h1>
        <p>Organize your trips, mark destinations as visited, edit travel notes, and plan future adventures.</p>
    </div>

    <!-- Alert Panel -->
    <div id="wishlist-alert" class="alert hidden"></div>

    <div class="card wishlist-card">
        <div class="table-responsive">
            <table class="wishlist-table">
                <thead>
                    <tr>
                        <th>Destination</th>
                        <th>Status</th>
                        <th>My Travel Notes</th>
                        <th>Added On</th>
                        <th class="actions-header">Actions</th>
                    </tr>
                </thead>
                <tbody id="wishlist-rows">
                    {{range .WishlistItems}}
                    <tr id="row-{{.ID}}" data-id="{{.ID}}">
                        <!-- Country Name -->
                        <td class="cell-destination">
                            <strong>{{.CountryName}}</strong>
                        </td>

                        <!-- Status Toggle (AJAX) -->
                        <td class="cell-status">
                            <select class="wishlist-status-select form-control select-sm" data-id="{{.ID}}">
                                <option value="Planned" {{if eq .Status "Planned" }}selected{{end}}>Planned</option>
                                <option value="Visited" {{if eq .Status "Visited" }}selected{{end}}>Visited</option>
                            </select>
                        </td>

                        <!-- Note Editing (AJAX) -->
                        <td class="cell-notes">
                            <div class="notes-edit-group">
                                <input type="text" class="wishlist-note-input form-control input-sm" data-id="{{.ID}}"
                                    value="{{.Note}}" placeholder="Add details or plans...">
                                <button type="button" class="btn btn-outline btn-sm btn-save-note" data-id="{{.ID}}">
                                    <i class="fa-solid fa-floppy-disk"></i>
                                </button>
                            </div>
                        </td>

                        <!-- Created Timestamp -->
                        <td class="cell-date">
                            {{.CreatedAt.Format "2006-01-02 15:04:05"}}
                        </td>

                        <!-- Delete Trigger (AJAX) -->
                        <td class="cell-actions">
                            <button type="button" class="btn btn-danger btn-sm btn-delete-wishlist" data-id="{{.ID}}">
                                <i class="fa-solid fa-trash-can"></i> Remove
                            </button>
                        </td>
                    </tr>
                    {{else}}
                    <tr class="empty-row">
                        <td colspan="5">
                            <div class="empty-wishlist">
                                <i class="fa-solid fa-heart-crack"></i>
                                <p>Your travel wishlist is currently empty.</p>
                                <a href="/countries" class="btn btn-primary"><i class="fa-solid fa-compass"></i>
                                    Discover Destinations</a>
                            </div>
                        </td>
                    </tr>
                    {{end}}
                </tbody>
            </table>
        </div>
    </div>
</div>