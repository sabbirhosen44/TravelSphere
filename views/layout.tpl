<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>TravelSphere | Discover the World</title>
    <!-- Google Fonts -->
    <link rel="preconnect" href="https://fonts.googleapis.com">
    <link rel="preconnect" href="https://fonts.gstatic.com" crossorigin>
    <link href="https://fonts.googleapis.com/css2?family=Plus+Jakarta+Sans:wght@300;400;500;600;700;800&display=swap"
        rel="stylesheet">
    <!-- FontAwesome for Icons -->
    <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/6.4.0/css/all.min.css">
    <!-- Premium Stylesheet -->
    <link rel="stylesheet" href="/static/css/main.css">
</head>

<body>
    <div class="app-container">
        <!-- Render Header -->
        {{.Header}}

        <!-- Main Content Wrapper -->
        <main class="main-content">
            {{if .ErrorMessage}}
            <div class="alert alert-danger" id="app-global-error">
                <i class="fa-solid fa-circle-exclamation"></i> {{.ErrorMessage}}
            </div>
            {{end}}

            {{if .SuccessMessage}}
            <div class="alert alert-success" id="app-global-success">
                <i class="fa-solid fa-circle-check"></i> {{.SuccessMessage}}
            </div>
            {{end}}

            {{.LayoutContent}}
        </main>

        <!-- Render Footer -->
        {{.Footer}}
    </div>

    <!-- Application AJAX Engine -->
    <script src="/static/js/ajax.js"></script>
</body>

</html>