function loadSnippet(snippetName) {
    fetch(`/snippet/${snippetName}`) // AJAX požadavek na server pro snippet
        .then(response => response.json()) // Očekáváme JSON odpověď
        .then(data => {
            const snippetId = "snippet-" + data.snippet; // ID snippetu (např. "content")
            const payload = data.payload;   // HTML obsah pro tento snippet

            // Vytvoření dynamického ID pro každý snippet
            const snippetDiv = document.querySelector(`#${snippetId}`);
            if (snippetDiv) {
                snippetDiv.innerHTML = payload;  // Dynamická změna obsahu
            }
        })
        .catch(error => console.error("Chyba při načítání snippetu:", error));
}