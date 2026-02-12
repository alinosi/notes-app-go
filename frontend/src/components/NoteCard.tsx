import React from 'react';
import { Edit, Trash2 } from 'lucide-react';
import type { Note } from '../types';
import { Button } from './Button';

interface NoteCardProps {
  note: Note;
  onEdit: (note: Note) => void;
  onDelete: (note: Note) => void;
}

export const NoteCard: React.FC<NoteCardProps> = ({ note, onEdit, onDelete }) => {
  const formatDate = (dateString: string) => {
    const date = new Date(dateString);
    const now = new Date();
    const diffInMs = now.getTime() - date.getTime();
    const diffInHours = diffInMs / (1000 * 60 * 60);

    if (diffInHours < 24) {
      return date.toLocaleTimeString('en-US', { hour: '2-digit', minute: '2-digit' });
    } else if (diffInHours < 168) {
      return date.toLocaleDateString('en-US', { weekday: 'short' });
    } else {
      return date.toLocaleDateString('en-US', { month: 'short', day: 'numeric' });
    }
  };

  const getPreview = (content: string, maxLength: number = 150) => {
    if (content.length <= maxLength) return content;
    return content.substring(0, maxLength) + '...';
  };

  return (
    <div className="bg-white border border-gray-200 rounded-lg p-4 hover:shadow-md transition-shadow">
      <div className="mb-3">
        <h3 className="font-semibold text-gray-900 mb-1 line-clamp-2">{note.title}</h3>
        <p className="text-sm text-gray-600 line-clamp-3">{getPreview(note.content)}</p>
      </div>
      <div className="flex items-center justify-between">
        <span className="text-xs text-gray-500">
          Updated {formatDate(note.updatedAt)}
        </span>
        <div className="flex gap-2">
          <Button
            variant="ghost"
            onClick={() => onEdit(note)}
            className="p-2"
            title="Edit note"
          >
            <Edit className="w-4 h-4" />
          </Button>
          <Button
            variant="ghost"
            onClick={() => onDelete(note)}
            className="p-2 text-red-600 hover:bg-red-50"
            title="Delete note"
          >
            <Trash2 className="w-4 h-4" />
          </Button>
        </div>
      </div>
    </div>
  );
};
