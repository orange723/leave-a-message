const LIMIT = 10;
let currentPage = 1;
let loading = false;
let lastPostedText = null;
let lastPostedAt = 0; // ms

async function fetchMessages(page = 1) {
  try {
    const res = await fetch(`/api/v1/message?limit=${LIMIT}&page=${page}`);
    if (!res.ok) throw new Error('Failed to load');
    return await res.json();
  } catch (err) {
    console.error(err);
    return [];
  }
}

function appendMessages(list) {
  const el = document.getElementById('messagesList');
  const empty = document.getElementById('empty');
  if (!list || list.length === 0) {
    if (el.children.length === 0) {
      empty.style.display = 'block';
    }
    return;
  }
  empty.style.display = 'none';
  list.forEach(m => {
    const li = document.createElement('li');
    const p = document.createElement('div');
    p.textContent = m.message;
    const t = document.createElement('time');
    const dt = new Date(m.createdAt || m.createdAt);
    t.textContent = dt.toLocaleString();
    li.appendChild(p);
    li.appendChild(t);
    el.appendChild(li);
  });
}

function clearMessages() {
  const el = document.getElementById('messagesList');
  el.innerHTML = '';
  document.getElementById('empty').style.display = 'none';
}

async function loadMore() {
  if (loading) return;
  loading = true;
  const btn = document.getElementById('loadMoreBtn');
  btn.disabled = true;
  btn.textContent = 'Loading...';
  const msgs = await fetchMessages(currentPage);
  appendMessages(msgs);
  if (!msgs || msgs.length < LIMIT) {
    btn.style.display = 'none';
  } else {
    btn.style.display = 'inline-block';
    btn.textContent = 'Load more';
  }
  if (msgs && msgs.length > 0) currentPage += 1;
  btn.disabled = false;
  loading = false;
}

async function postMessage(text) {
  // client-side de-dup: avoid posting same text repeatedly within 5s
  const now = Date.now();
  if (text === lastPostedText && (now - lastPostedAt) < 5000) {
    return true;
  }
  try {
    const res = await fetch('/api/v1/message', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ message: text }),
    });
    if (res.ok) {
      lastPostedText = text;
      lastPostedAt = Date.now();
    }
    return res.ok;
  } catch (err) { console.error(err); return false; }
}

document.addEventListener('DOMContentLoaded', () => {
  // initial load
  loadMore();

  const form = document.getElementById('msgForm');
  const textarea = document.getElementById('message');
  form.addEventListener('submit', async (e) => {
    e.preventDefault();
    const v = textarea.value.trim();
    if (!v) return;
    const submitBtn = form.querySelector('button[type="submit"]');
    if (submitBtn) submitBtn.disabled = true;
    const ok = await postMessage(v);
    if (ok) {
      textarea.value = '';
      // reset and reload first page
      currentPage = 1;
      clearMessages();
      document.getElementById('loadMoreBtn').style.display = 'inline-block';
      // replace history state to avoid browser re-submitting on refresh
      try { history.replaceState(null, '', window.location.pathname); } catch (e) {}
      loadMore();
    } else {
      alert('Failed to post message');
    }
    if (submitBtn) submitBtn.disabled = false;
  });

  const loadBtn = document.getElementById('loadMoreBtn');
  loadBtn.addEventListener('click', async () => {
    await loadMore();
  });
});
