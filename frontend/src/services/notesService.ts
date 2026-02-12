import type { Note, CreateNoteData, UpdateNoteData } from '../types';
import { storage } from '../utils/storage';

// Simulate network delay
const delay = (ms: number) => new Promise(resolve => setTimeout(resolve, ms));

// Mock notes service using localStorage
export const notesService = {
  async getNotes(searchQuery?: string): Promise<Note[]> {
    await delay(600);

    const token = storage.getToken();
    if (!token) {
      throw new Error('Unauthorized');
    }

    // Extract user ID from token
    const userIdMatch = token.match(/mock_token_([^_]+)/);
    if (!userIdMatch) {
      throw new Error('Invalid token');
    }

    const userId = userIdMatch[1];
    const allNotes = storage.getNotes();
    
    // Filter notes by user
    let userNotes = allNotes.filter((note: Note) => note.userId === userId);

    // Apply search filter if provided
    if (searchQuery && searchQuery.trim()) {
      const query = searchQuery.toLowerCase();
      userNotes = userNotes.filter(
        (note: Note) =>
          note.title.toLowerCase().includes(query) ||
          note.content.toLowerCase().includes(query)
      );
    }

    // Sort by updated date (most recent first)
    userNotes.sort((a: Note, b: Note) => 
      new Date(b.updatedAt).getTime() - new Date(a.updatedAt).getTime()
    );

    return userNotes;
  },

  async getNoteById(id: string): Promise<Note> {
    await delay(400);

    const token = storage.getToken();
    if (!token) {
      throw new Error('Unauthorized');
    }

    const allNotes = storage.getNotes();
    const note = allNotes.find((n: Note) => n.id === id);

    if (!note) {
      throw new Error('Note not found');
    }

    // Verify note belongs to current user
    const userIdMatch = token.match(/mock_token_([^_]+)/);
    if (!userIdMatch || note.userId !== userIdMatch[1]) {
      throw new Error('Unauthorized');
    }

    return note;
  },

  async createNote(data: CreateNoteData): Promise<Note> {
    await delay(500);

    const token = storage.getToken();
    if (!token) {
      throw new Error('Unauthorized');
    }

    const userIdMatch = token.match(/mock_token_([^_]+)/);
    if (!userIdMatch) {
      throw new Error('Invalid token');
    }

    const userId = userIdMatch[1];
    const now = new Date().toISOString();

    const newNote: Note = {
      id: `note_${Date.now()}`,
      title: data.title,
      content: data.content,
      userId,
      createdAt: now,
      updatedAt: now,
    };

    const allNotes = storage.getNotes();
    allNotes.push(newNote);
    storage.setNotes(allNotes);

    return newNote;
  },

  async updateNote(id: string, data: UpdateNoteData): Promise<Note> {
    await delay(500);

    const token = storage.getToken();
    if (!token) {
      throw new Error('Unauthorized');
    }

    const userIdMatch = token.match(/mock_token_([^_]+)/);
    if (!userIdMatch) {
      throw new Error('Invalid token');
    }

    const userId = userIdMatch[1];
    const allNotes = storage.getNotes();
    const noteIndex = allNotes.findIndex((n: Note) => n.id === id);

    if (noteIndex === -1) {
      throw new Error('Note not found');
    }

    if (allNotes[noteIndex].userId !== userId) {
      throw new Error('Unauthorized');
    }

    const updatedNote: Note = {
      ...allNotes[noteIndex],
      title: data.title,
      content: data.content,
      updatedAt: new Date().toISOString(),
    };

    allNotes[noteIndex] = updatedNote;
    storage.setNotes(allNotes);

    return updatedNote;
  },

  async deleteNote(id: string): Promise<void> {
    await delay(400);

    const token = storage.getToken();
    if (!token) {
      throw new Error('Unauthorized');
    }

    const userIdMatch = token.match(/mock_token_([^_]+)/);
    if (!userIdMatch) {
      throw new Error('Invalid token');
    }

    const userId = userIdMatch[1];
    const allNotes = storage.getNotes();
    const noteIndex = allNotes.findIndex((n: Note) => n.id === id);

    if (noteIndex === -1) {
      throw new Error('Note not found');
    }

    if (allNotes[noteIndex].userId !== userId) {
      throw new Error('Unauthorized');
    }

    allNotes.splice(noteIndex, 1);
    storage.setNotes(allNotes);
  },
};
