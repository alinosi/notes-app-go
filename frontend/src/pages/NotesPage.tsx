import React, { useState, useEffect } from 'react';
import { useNavigate } from 'react-router';
import { Plus, AlertCircle, FileText } from 'lucide-react';
import { AppLayout } from '../layouts/AppLayout';
import { Button } from '../components/Button';
import { NoteCard } from '../components/NoteCard';
import { NoteCardSkeleton } from '../components/NoteCardSkeleton';
import { Modal } from '../components/Modal';
import { notesService } from '../services/notesService';
import type { Note } from '../types';
import { toast } from 'sonner@2.0.3';

export const NotesPage: React.FC = () => {
  const navigate = useNavigate();
  const [notes, setNotes] = useState<Note[]>([]);
  const [isLoading, setIsLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);
  const [searchQuery, setSearchQuery] = useState('');
  const [noteToDelete, setNoteToDelete] = useState<Note | null>(null);
  const [isDeleting, setIsDeleting] = useState(false);

  const fetchNotes = async (query?: string) => {
    setIsLoading(true);
    setError(null);
    try {
      const data = await notesService.getNotes(query);
      setNotes(data);
    } catch (err: any) {
      setError(err.message || 'Failed to load notes');
    } finally {
      setIsLoading(false);
    }
  };

  useEffect(() => {
    fetchNotes();
  }, []);

  const handleSearch = (query: string) => {
    setSearchQuery(query);
    fetchNotes(query);
  };

  const handleEdit = (note: Note) => {
    navigate(`/notes/${note.id}/edit`);
  };

  const handleDeleteClick = (note: Note) => {
    setNoteToDelete(note);
  };

  const handleDeleteConfirm = async () => {
    if (!noteToDelete) return;

    setIsDeleting(true);
    try {
      await notesService.deleteNote(noteToDelete.id);
      setNotes(prev => prev.filter(n => n.id !== noteToDelete.id));
      toast.success('Note deleted');
      setNoteToDelete(null);
    } catch (error: any) {
      toast.error(error.message || 'Failed to delete note');
    } finally {
      setIsDeleting(false);
    }
  };

  const handleRetry = () => {
    fetchNotes(searchQuery);
  };

  return (
    <AppLayout onSearch={handleSearch}>
      <div className="mb-8 flex items-center justify-between">
        <h1 className="text-3xl font-bold text-gray-900">My Notes</h1>
        <Button
          variant="primary"
          onClick={() => navigate('/notes/new')}
          className="flex items-center gap-2"
        >
          <Plus className="w-5 h-5" />
          New Note
        </Button>
      </div>

      {/* Loading State */}
      {isLoading && (
        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
          {[1, 2, 3, 4, 5, 6].map(i => (
            <NoteCardSkeleton key={i} />
          ))}
        </div>
      )}

      {/* Error State */}
      {error && !isLoading && (
        <div className="bg-red-50 border border-red-200 rounded-lg p-8 text-center">
          <AlertCircle className="w-12 h-12 text-red-600 mx-auto mb-4" />
          <h3 className="text-lg font-semibold text-gray-900 mb-2">Failed to Load Notes</h3>
          <p className="text-gray-600 mb-4">{error}</p>
          <Button variant="primary" onClick={handleRetry}>
            Try Again
          </Button>
        </div>
      )}

      {/* Empty State */}
      {!isLoading && !error && notes.length === 0 && (
        <div className="bg-white border border-gray-200 rounded-lg p-12 text-center">
          <FileText className="w-16 h-16 text-gray-400 mx-auto mb-4" />
          <h3 className="text-xl font-semibold text-gray-900 mb-2">
            {searchQuery ? 'No notes found' : 'No notes yet'}
          </h3>
          <p className="text-gray-600 mb-6">
            {searchQuery
              ? 'Try adjusting your search query'
              : 'Get started by creating your first note'}
          </p>
          {!searchQuery && (
            <Button
              variant="primary"
              onClick={() => navigate('/notes/new')}
              className="inline-flex items-center gap-2"
            >
              <Plus className="w-5 h-5" />
              Create Your First Note
            </Button>
          )}
        </div>
      )}

      {/* Notes Grid */}
      {!isLoading && !error && notes.length > 0 && (
        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
          {notes.map(note => (
            <NoteCard
              key={note.id}
              note={note}
              onEdit={handleEdit}
              onDelete={handleDeleteClick}
            />
          ))}
        </div>
      )}

      {/* Delete Confirmation Modal */}
      <Modal
        isOpen={!!noteToDelete}
        onClose={() => !isDeleting && setNoteToDelete(null)}
        title="Delete Note?"
      >
        <div className="space-y-4">
          <p className="text-gray-600">
            Are you sure you want to delete <strong>"{noteToDelete?.title}"</strong>?
            This action cannot be undone.
          </p>
          <div className="flex gap-3 justify-end">
            <Button
              variant="secondary"
              onClick={() => setNoteToDelete(null)}
              disabled={isDeleting}
            >
              Cancel
            </Button>
            <Button
              variant="danger"
              onClick={handleDeleteConfirm}
              isLoading={isDeleting}
            >
              Delete
            </Button>
          </div>
        </div>
      </Modal>
    </AppLayout>
  );
};
