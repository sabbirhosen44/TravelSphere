document.addEventListener('DOMContentLoaded', () => {

    // HOME PAGE AUTOCOMPLETE SEARCH SUGGESTIONS

    const homeSearch = document.getElementById('home-search');
    const homeSpinner = document.getElementById('home-search-spinner');
    const homeResults = document.getElementById('home-search-results');

    if (homeSearch && homeResults) {
        let debounceTimer;
        homeSearch.addEventListener('input', () => {
            clearTimeout(debounceTimer);
            const val = homeSearch.value.trim();
            if (val === '') {
                homeResults.innerHTML = '';
                homeResults.classList.add('hidden');
                return;
            }

            if (homeSpinner) homeSpinner.classList.remove('hidden');

            debounceTimer = setTimeout(() => {
                fetch(`/api/countries?search=${encodeURIComponent(val)}`)
                    .then(res => res.json())
                    .then(data => {
                        if (homeSpinner) homeSpinner.classList.add('hidden');
                        if (data && data.length > 0) {
                            homeResults.innerHTML = '';
                            homeResults.classList.remove('hidden');
                            data.slice(0, 5).forEach(country => {
                                const div = document.createElement('div');
                                div.className = 'dropdown-item';
                                div.innerHTML = `
                                    <img src="${country.flag}" alt="" class="flag-icon-mini" style="height: 16px; width: auto; vertical-align: middle; margin-right: 8px; border-radius: 2px;"> 
                                    <span>${country.common_name}</span> 
                                    <small class="text-muted">(${country.capital})</small>
                                `;
                                div.addEventListener('click', () => {
                                    window.location.href = `/countries/${country.slug}`;
                                });
                                homeResults.appendChild(div);
                            });
                        } else {
                            homeResults.innerHTML = '<div class="dropdown-item no-results-item">No destinations found</div>';
                            homeResults.classList.remove('hidden');
                        }
                    })
                    .catch(() => {
                        if (homeSpinner) homeSpinner.classList.add('hidden');
                    });
            }, 300);
        });

        document.addEventListener('click', (e) => {
            if (e.target !== homeSearch && e.target !== homeResults) {
                homeResults.classList.add('hidden');
            }
        });
    }


    // COUNTRY EXPLORER (LIVE FILTER, SEARCH & PAGINATION)

    const explorerSearch = document.getElementById('explorer-search');
    const explorerRegion = document.getElementById('explorer-region');
    const explorerSpinner = document.getElementById('explorer-spinner');
    const countryResults = document.getElementById('country-results');
    const paginationContainer = document.getElementById('pagination-container');

    if (explorerSearch || explorerRegion) {
        const fetchExplorerData = (page = 1) => {
            const searchVal = explorerSearch ? explorerSearch.value.trim() : '';
            const regionVal = explorerRegion ? explorerRegion.value : '';

            if (explorerSpinner) explorerSpinner.classList.remove('hidden');

            fetch(`/api/countries?search=${encodeURIComponent(searchVal)}&region=${encodeURIComponent(regionVal)}&page=${page}`)
                .then(res => res.json())
                .then(data => {
                    if (explorerSpinner) explorerSpinner.classList.add('hidden');

                    if (!countryResults) return;
                    countryResults.innerHTML = '';

                    let parsedCountries = [];
                    let totalPages = 1;
                    let currentPage = 1;
                    let hasPrev = false;
                    let hasNext = false;
                    let prevPage = 1;
                    let nextPage = 1;

                    if (Array.isArray(data)) {
                        parsedCountries = data;
                    } else if (data && data.countries) {
                        parsedCountries = data.countries;
                        totalPages = data.total_pages || 1;
                        currentPage = data.current_page || 1;
                        hasPrev = data.has_prev || false;
                        hasNext = data.has_next || false;
                        prevPage = data.prev_page || 1;
                        nextPage = data.next_page || 1;
                    }

                    if (parsedCountries && parsedCountries.length > 0) {
                        parsedCountries.forEach(c => {
                            const card = document.createElement('div');
                            card.className = 'country-card';
                            card.innerHTML = `
                                <div class="country-card-header">
                                    <img src="${c.flag}" alt="${c.common_name} flag" class="country-flag-large" style="height: 40px; width: auto; object-fit: cover; border-radius: 4px;">
                                    <span class="country-region-badge">${c.region}</span>
                                </div>
                                <div class="country-card-body">
                                    <h3>${c.common_name}</h3>
                                    <p><strong>Capital:</strong> ${c.capital}</p>
                                    <p><strong>Population:</strong> ${c.formatted_population}</p>
                                    <p><strong>Currency:</strong> ${c.currencies || 'N/A'}</p>
                                    <p><strong>Languages:</strong> ${c.languages || 'N/A'}</p>
                                </div>
                                <div class="country-card-footer">
                                    <a href="/countries/${c.slug}" class="btn btn-primary btn-block">Explore Destination</a>
                                </div>
                            `;
                            countryResults.appendChild(card);
                        });

                        if (paginationContainer) {
                            paginationContainer.innerHTML = '';
                            if (totalPages > 1) {
                                const nav = document.createElement('nav');
                                nav.className = 'pagination-nav';

                                let prevLink = '';
                                if (hasPrev) {
                                    prevLink = `<button class="btn-pagination prev-page" data-page="${prevPage}"><i class="fa-solid fa-chevron-left"></i> Previous</button>`;
                                }

                                let nextLink = '';
                                if (hasNext) {
                                    nextLink = `<button class="btn-pagination next-page" data-page="${nextPage}">Next <i class="fa-solid fa-chevron-right"></i></button>`;
                                }

                                nav.innerHTML = `
                                    ${prevLink}
                                    <span class="page-info">Page <strong>${currentPage}</strong> of <strong>${totalPages}</strong></span>
                                    ${nextLink}
                                `;
                                paginationContainer.appendChild(nav);
                            }
                        }
                    } else {
                        countryResults.innerHTML = `
                            <div class="no-results">
                                <i class="fa-solid fa-earth-europe"></i>
                                <p>No countries matching filters were found.</p>
                            </div>
                        `;
                        if (paginationContainer) paginationContainer.innerHTML = '';
                    }
                })
                .catch(err => {
                    if (explorerSpinner) explorerSpinner.classList.add('hidden');
                    if (countryResults) {
                        countryResults.innerHTML = `<div class="alert alert-danger">Error fetching destinations: ${err.message}</div>`;
                    }
                });
        };

        let searchDebounce;
        if (explorerSearch) {
            explorerSearch.addEventListener('input', () => {
                clearTimeout(searchDebounce);
                searchDebounce = setTimeout(() => fetchExplorerData(1), 300);
            });
        }

        if (explorerRegion) {
            explorerRegion.addEventListener('change', () => {
                fetchExplorerData(1);
            });
        }

        if (paginationContainer) {
            paginationContainer.addEventListener('click', (e) => {
                const btn = e.target.closest('.btn-pagination');
                if (btn) {
                    e.preventDefault();
                    const targetPage = parseInt(btn.getAttribute('data-page'), 10);
                    if (targetPage) {
                        fetchExplorerData(targetPage);
                        if (explorerSearch) {
                            window.scrollTo({ top: explorerSearch.offsetTop - 20, behavior: 'smooth' });
                        }
                    }
                }
            });
        }
    }


    // DESTINATION DETAILS (ADD TO WISHLIST)

    const btnAddWishlist = document.getElementById('btn-add-wishlist');
    const wishlistFeedback = document.getElementById('wishlist-feedback');

    if (btnAddWishlist && wishlistFeedback) {
        btnAddWishlist.addEventListener('click', () => {
            const countryName = btnAddWishlist.getAttribute('data-country');
            const status = document.getElementById('wishlist-status').value;
            const note = document.getElementById('wishlist-note').value;

            btnAddWishlist.disabled = true;
            btnAddWishlist.innerHTML = '<i class="fa-solid fa-spinner fa-spin"></i> Adding...';

            fetch('/api/wishlist', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify({
                    country_name: countryName,
                    status: status,
                    note: note
                })
            })
                .then(res => {
                    if (!res.ok) {
                        return res.json().then(data => { throw new Error(data.error || 'Failed to add item') });
                    }
                    return res.json();
                })
                .then(() => {
                    wishlistFeedback.innerHTML = `
                        <div class="alert alert-info">
                            <p><i class="fa-solid fa-circle-check"></i> Added to your wishlist!</p>
                            <a href="/wishlist" class="btn btn-outline btn-sm"><i class="fa-solid fa-heart"></i> Manage Wishlist</a>
                        </div>
                    `;
                })
                .catch(err => {
                    btnAddWishlist.disabled = false;
                    btnAddWishlist.innerHTML = '<i class="fa-solid fa-circle-plus"></i> Add Destination to Wishlist';

                    const errDiv = document.createElement('div');
                    errDiv.className = 'alert alert-danger mt-2';
                    errDiv.innerHTML = `<i class="fa-solid fa-circle-exclamation"></i> ${err.message}`;

                    const existingAlert = wishlistFeedback.querySelector('.alert-danger');
                    if (existingAlert) existingAlert.remove();

                    wishlistFeedback.appendChild(errDiv);
                });
        });
    }


    // WISHLIST MANAGEMENT (INLINE UPDATE & DELETE)

    const wishlistRows = document.getElementById('wishlist-rows');
    const wishlistAlert = document.getElementById('wishlist-alert');

    const showWishlistAlert = (msg, isSuccess) => {
        if (!wishlistAlert) return;
        wishlistAlert.className = `alert ${isSuccess ? 'alert-success' : 'alert-danger'}`;
        wishlistAlert.innerHTML = `<i class="fa-solid ${isSuccess ? 'fa-circle-check' : 'fa-circle-exclamation'}"></i> ${msg}`;
        wishlistAlert.classList.remove('hidden');
        setTimeout(() => {
            wishlistAlert.classList.add('hidden');
        }, 3000);
    };

    if (wishlistRows) {
        wishlistRows.addEventListener('click', (e) => {
            const btnSave = e.target.closest('.btn-save-note');
            if (btnSave) {
                const id = btnSave.getAttribute('data-id');
                const row = document.getElementById(`row-${id}`);
                const noteVal = row.querySelector('.wishlist-note-input').value;
                const statusVal = row.querySelector('.wishlist-status-select').value;

                btnSave.disabled = true;
                btnSave.innerHTML = '<i class="fa-solid fa-spinner fa-spin"></i>';

                fetch(`/api/wishlist/${id}`, {
                    method: 'PUT',
                    headers: {
                        'Content-Type': 'application/json'
                    },
                    body: JSON.stringify({
                        status: statusVal,
                        note: noteVal
                    })
                })
                    .then(res => {
                        if (!res.ok) {
                            return res.json().then(data => { throw new Error(data.error || 'Update failed') });
                        }
                        return res.json();
                    })
                    .then(() => {
                        btnSave.disabled = false;
                        btnSave.innerHTML = '<i class="fa-solid fa-floppy-disk"></i>';
                        showWishlistAlert('Note saved successfully!', true);
                    })
                    .catch(err => {
                        btnSave.disabled = false;
                        btnSave.innerHTML = '<i class="fa-solid fa-floppy-disk"></i>';
                        showWishlistAlert(`Error saving note: ${err.message}`, false);
                    });
            }
        });

        wishlistRows.addEventListener('change', (e) => {
            const selectStatus = e.target.closest('.wishlist-status-select');
            if (selectStatus) {
                const id = selectStatus.getAttribute('data-id');
                const row = document.getElementById(`row-${id}`);
                const noteVal = row.querySelector('.wishlist-note-input').value;
                const statusVal = selectStatus.value;

                selectStatus.disabled = true;

                fetch(`/api/wishlist/${id}`, {
                    method: 'PUT',
                    headers: {
                        'Content-Type': 'application/json'
                    },
                    body: JSON.stringify({
                        status: statusVal,
                        note: noteVal
                    })
                })
                    .then(res => {
                        if (!res.ok) {
                            return res.json().then(data => { throw new Error(data.error || 'Status update failed') });
                        }
                        return res.json();
                    })
                    .then(() => {
                        selectStatus.disabled = false;
                        showWishlistAlert('Status updated successfully!', true);
                    })
                    .catch(err => {
                        selectStatus.disabled = false;
                        showWishlistAlert(`Error updating status: ${err.message}`, false);
                    });
            }
        });

        wishlistRows.addEventListener('click', (e) => {
            const btnDel = e.target.closest('.btn-delete-wishlist');
            if (btnDel) {
                const id = btnDel.getAttribute('data-id');
                if (!confirm('Are you sure you want to remove this destination?')) return;

                btnDel.disabled = true;
                btnDel.innerHTML = '<i class="fa-solid fa-spinner fa-spin"></i> Removing...';

                fetch(`/api/wishlist/${id}`, {
                    method: 'DELETE'
                })
                    .then(res => {
                        if (!res.ok) {
                            return res.json().then(data => { throw new Error(data.error || 'Delete failed') });
                        }
                        return res.json();
                    })
                    .then(() => {
                        const row = document.getElementById(`row-${id}`);
                        if (row) {
                            row.style.transition = 'opacity 0.3s ease';
                            row.style.opacity = '0';
                            setTimeout(() => {
                                row.remove();
                                if (wishlistRows.querySelectorAll('tr[id^="row-"]').length === 0) {
                                    wishlistRows.innerHTML = `
                                        <tr class="empty-row">
                                            <td colspan="5">
                                                <div class="empty-wishlist">
                                                    <i class="fa-solid fa-heart-crack"></i>
                                                    <p>Your travel wishlist is currently empty.</p>
                                                    <a href="/countries" class="btn btn-primary"><i class="fa-solid fa-compass"></i> Discover Destinations</a>
                                                </div>
                                            </td>
                                        </tr>
                                    `;
                                }
                            }, 300);
                        }
                        showWishlistAlert('Destination removed successfully!', true);
                    })
                    .catch(err => {
                        btnDel.disabled = false;
                        btnDel.innerHTML = '<i class="fa-solid fa-trash-can"></i> Remove';
                        showWishlistAlert(`Error removing item: ${err.message}`, false);
                    });
            }
        });
    }
});