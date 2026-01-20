// --- KONFIGURASI API ---
const API_URL = "http://localhost:8080/api/notes";
let currentDeleteId = null;

// --- IKON SVG (Disimpan dalam variabel agar kode render lebih rapi) ---
const deleteIconSVG = `<svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><polyline points="3 6 5 6 21 6"></polyline><path d="M19 6v14a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2V6m3 0V4a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2v2"></path><line x1="10" y1="11" x2="10" y2="17"></line><line x1="14" y1="11" x2="14" y2="17"></line></svg>`;

// --- FUNGSI UTAMA ---

document.addEventListener('DOMContentLoaded', () => {
    // Hapus data mock ini nanti dan gunakan fetchNotes() yang asli
    const mockNotes = [
        { id: 1, title: "Ide Proyek Golang", content: "Membuat backend REST API yang cepat menggunakan Gin atau standar library. Perlu memikirkan struktur database dan autentikasi JWT nanti.", created_at: "2023-10-25" },
        { id: 2, title: "Daftar Belanja Mingguan", content: "Telur, Susu, Roti gandum, Sayur bayam, Buah apel, Kopi.", created_at: "2023-10-24" },
        { id: 3, title: "Catatan Meeting UI/UX", content: "Fokus pada penggunaan whitespace yang lebih banyak. Warna harus lebih lembut. Animasi tidak boleh berlebihan, harus terasa natural.", created_at: "2023-10-23" },
         { id: 4, title: "Jadwal Olahraga", content: "Senin: Lari 5km. Rabu: Angkat beban. Jumat: Berenang.", created_at: "2023-10-22" }
    ];
    renderNotes(mockNotes); // Menggunakan data palsu untuk preview UI

    // UNCOMMENT INI UNTUK MENGGUNAKAN BACKEND GOLANG ASLI:
    // fetchNotes();
});


// 1. Mengambil Data dari API
async function fetchNotes() {
    try {
        const response = await fetch(API_URL);
        if (!response.ok) throw new Error('Gagal mengambil data');
        const notes = await response.json();
        renderNotes(notes);
    } catch (error) {
        console.error("Error:", error);
        // Opsional: Tampilkan pesan error di UI
        document.getElementById('notes-grid').innerHTML = `<p style="color: red; text-align:center;">Gagal memuat catatan. Pastikan backend Golang berjalan.</p>`;
    }
}

// 2. Render Kartu Catatan ke HTML
function renderNotes(notes) {
    const grid = document.getElementById('notes-grid');
    const emptyState = document.getElementById('empty-state');
    grid.innerHTML = ''; // Bersihkan grid

    if (!notes || notes.length === 0) {
        emptyState.classList.remove('hidden');
        grid.classList.add('hidden');
        return;
    }

    emptyState.classList.add('hidden');
    grid.classList.remove('hidden');

    notes.forEach(note => {
        const card = document.createElement('div');
        card.className = 'note-card';
        
        // Format tanggal sederhana (opsional, sesuaikan dengan format dari Go)
        const dateStr = note.created_at ? new Date(note.created_at).toLocaleDateString('id-ID', { year: 'numeric', month: 'short', day: 'numeric' }) : '';

        card.innerHTML = `
            <div class="card-content" onclick="viewDetail(${note.id})">
                <h3>${note.title}</h3>
                <p>${note.content}</p>
                <span class="card-date">${dateStr}</span>
            </div>
            <button class="btn-delete-action" onclick="openDeleteModal(event, ${note.id})">
                ${deleteIconSVG}
            </button>
        `;
        grid.appendChild(card);
    });
}

// 3. Menampilkan Detail Catatan
async function viewDetail(id) {
    // Untuk preview UI dengan mock data (HAPUS BLOK INI SAAT INTEGRASI)
    const mockNote = { id: id, title: "Detail Catatan (Mock)", content: "Ini adalah konten lengkap dari catatan yang Anda klik. Di sini Anda bisa membaca seluruh teks tanpa terpotong.\n\nBaris baru juga akan terlihat rapi di sini.", created_at: "2023-10-25" };
    renderDetailView(mockNote);
    return; 
    // ------------------------------------------------------------------

    // UNCOMMENT UNTUK INTEGRASI ASLI:
    /*
    try {
        const response = await fetch(`${API_URL}/${id}`);
        if (!response.ok) throw new Error('Catatan tidak ditemukan');
        const note = await response.json();
        renderDetailView(note);
    } catch (error) {
        alert("Gagal memuat detail catatan: " + error.message);
    }
    */
}

// Fungsi helper untuk merender tampilan detail
function renderDetailView(note) {
    document.getElementById('notes-list-view').classList.add('hidden');
    const detailView = document.getElementById('note-detail-view');
    detailView.classList.remove('hidden');
    
    // Scroll ke atas saat membuka detail
    window.scrollTo(0,0);

    const dateStr = note.created_at ? new Date(note.created_at).toLocaleDateString('id-ID', { weekday:'long', year: 'numeric', month: 'long', day: 'numeric' }) : '';

    document.getElementById('detail-content').innerHTML = `
        <h2 class="detail-title">${note.title}</h2>
        <div class="detail-body">${note.content}</div>
        <div class="detail-meta">Dibuat pada: ${dateStr}</div>
    `;
}

function showListView() {
    document.getElementById('notes-list-view').classList.remove('hidden');
    document.getElementById('note-detail-view').classList.add('hidden');
}

// --- Fungsi Modal Tambah ---
function openAddModal() {
    document.getElementById('add-modal').classList.remove('hidden');
}

function closeAddModal() {
    document.getElementById('add-modal').classList.add('hidden');
    document.getElementById('note-form').reset(); // Bersihkan form
}

// --- Handle Submit Form ---
document.getElementById('note-form').addEventListener('submit', async (e) => {
    e.preventDefault();

    const title = document.getElementById('note-title').value;
    const content = document.getElementById('note-content').value;

    const newNote = {
        title: title,
        content: content
    };

    try {
        // HTTP Request ke Golang Backend
        const response = await fetch(API_URL, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(newNote)
        });

        if (response.ok) {
            closeAddModal();
            fetchNotes(); // Refresh daftar catatan
        } else {
            alert("Gagal menyimpan catatan");
        }
    } catch (error) {
        console.error("Error:", error);
        // Simulasi untuk kebutuhan preview tanpa backend:
        alert("Catatan Berhasil Disimpan (Mode Simulasi)");
        closeAddModal();
    }
});

// 4. Logika Modal Hapus
function openDeleteModal(event, id) {
    event.stopPropagation(); // Mencegah klik tembus ke card (mencegah viewDetail)
    currentDeleteId = id;
    document.getElementById('delete-modal').classList.remove('hidden');
}

function closeModal() {
    document.getElementById('delete-modal').classList.add('hidden');
    currentDeleteId = null;
}

// Event listener untuk tombol konfirmasi di modal
document.getElementById('confirm-delete-btn').addEventListener('click', async () => {
    if (currentDeleteId) {
        // UNCOMMENT UNTUK INTEGRASI ASLI:
        /*
        try {
            const response = await fetch(`${API_URL}/${currentDeleteId}`, { method: 'DELETE' });
            if (!response.ok) throw new Error('Gagal menghapus');
        } catch (error) {
             alert("Error saat menghapus: " + error.message);
        }
        */

        console.log(`Simulasi menghapus catatan ID: ${currentDeleteId}`); // Hapus baris ini nanti
        closeModal();
        // fetchNotes(); // Refresh list setelah hapus (uncomment nanti)
        alert("Catatan berhasil dihapus (Simulasi). Refresh halaman untuk reset mock data."); // Hapus alert ini nanti
    }
});

// Menutup modal jika mengklik area gelap di luar box modal
document.getElementById('delete-modal').addEventListener('click', (e) => {
    if (e.target === document.getElementById('delete-modal')) {
        closeModal();
    }
});