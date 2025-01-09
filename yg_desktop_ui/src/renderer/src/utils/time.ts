export function formatTimestamp(timestamp: string, locale: string = 'zh-CN'): string {
  const date = new Date(timestamp);
  return new Intl.DateTimeFormat(locale, {
    hour: '2-digit',
    minute: '2-digit',
    second: '2-digit',
    hour12: true
  }).format(date);
}


export function formatLocalizedSessionTime(timestamp: string): string {
  const date = new Date(timestamp);
  const now = new Date();
  const isToday = date.toDateString() === now.toDateString();
  const isYesterday = new Date(now.getTime() - 24 * 60 * 60 * 1000).toDateString() === date.toDateString();
  const isThisWeek = now.getTime() - date.getTime() < 7 * 24 * 60 * 60 * 1000 && now.getDay() >= date.getDay();
  const isThisYear = now.getFullYear() === date.getFullYear();

  if (isToday) {
    return new Intl.DateTimeFormat('zh-CN', {
      hour: '2-digit',
      minute: '2-digit'
    }).format(date);
  } else if (isYesterday) {
    return `昨天 ${new Intl.DateTimeFormat('zh-CN', {
      hour: '2-digit',
      minute: '2-digit'
    }).format(date)}`;
  } else if (isThisWeek) {
    return `${new Intl.DateTimeFormat('zh-CN', {
      weekday: 'long',
      hour: '2-digit',
      minute: '2-digit'
    }).format(date)}`;
  } else if (isThisYear) {
    return new Intl.DateTimeFormat('zh-CN', {
      month: '2-digit',
      day: '2-digit',
      hour: '2-digit',
      minute: '2-digit'
    }).format(date);
  } else {
    return new Intl.DateTimeFormat('zh-CN', {
      year: 'numeric',
      month: '2-digit',
      day: '2-digit',
      hour: '2-digit',
      minute: '2-digit'
    }).format(date);
  }
}