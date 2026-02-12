// Helper functions for localStorage management

export const storage = {
  getToken: (): string | null => {
    return localStorage.getItem('auth_token');
  },

  setToken: (token: string): void => {
    localStorage.setItem('auth_token', token);
  },

  removeToken: (): void => {
    localStorage.removeItem('auth_token');
  },

  getUsers: (): any[] => {
    const users = localStorage.getItem('users');
    return users ? JSON.parse(users) : [];
  },

  setUsers: (users: any[]): void => {
    localStorage.setItem('users', JSON.stringify(users));
  },

  getNotes: (): any[] => {
    const notes = localStorage.getItem('notes');
    return notes ? JSON.parse(notes) : [];
  },

  setNotes: (notes: any[]): void => {
    localStorage.setItem('notes', JSON.stringify(notes));
  },
};
