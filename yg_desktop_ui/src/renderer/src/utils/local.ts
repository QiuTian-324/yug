export function getLocalCurrentSessionId() {
  const chatStoreData = localStorage.getItem('chatStore') || '{}';
  return JSON.parse(chatStoreData).selectedSessionId;
}