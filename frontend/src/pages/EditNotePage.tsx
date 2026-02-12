import React, { useState, useEffect } from 'react';
import { useNavigate, useParams } from 'react-router';
import { ArrowLeft } from 'lucide-react';
import { AppLayout } from '../layouts/AppLayout';
import { Input } from '../components/Input';
import { Textarea } from '../components/Textarea';
import { Button } from '../components/Button';
import { notesService } from '../services/notesService';
import { toast } from 'sonner@2.0.3';

export const EditNotePage: React.FC = () => {
  const navigate = useNavigate();
  const { id } = useParams<{ id: string }>();
  
  const [formData, setFormData] = useState({
    title: '',
    content: '',
  });
  const [errors, setErrors] = useState<Record<string, string>>({});
  const [isLoading, setIsLoading] = useState(true);
  const [isSaving, setIsSaving] = useState(false);

  useEffect(() => {
    const loadNote = async () => {
      if (!id) {
        navigate('/notes');
        return;
      }

      try {
        const note = await notesService.getNoteById(id);
        setFormData({
          title: note.title,
          content: note.content,
        });
      } catch (error: any) {
        toast.error(error.message || 'Failed to load note');
        navigate('/notes');
      } finally {
        setIsLoading(false);
      }
    };

    loadNote();
  }, [id, navigate]);

  const handleChange = (e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>) => {
    const { name, value } = e.target;
    setFormData(prev => ({ ...prev, [name]: value }));
    
    if (errors[name]) {
      setErrors(prev => ({ ...prev, [name]: undefined }));
    }
  };

  const validate = (): boolean => {
    const newErrors: Record<string, string> = {};

    if (!formData.title.trim()) {
      newErrors.title = 'Title is required';
    }

    if (!formData.content.trim()) {
      newErrors.content = 'Content is required';
    }

    setErrors(newErrors);
    return Object.keys(newErrors).length === 0;
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();

    if (!validate() || !id) return;

    setIsSaving(true);

    try {
      await notesService.updateNote(id, {
        title: formData.title,
        content: formData.content,
      });

      toast.success('Note updated');
      navigate('/notes');
    } catch (error: any) {
      toast.error(error.message || 'Failed to update note');
    } finally {
      setIsSaving(false);
    }
  };

  const handleCancel = () => {
    navigate('/notes');
  };

  if (isLoading) {
    return (
      <AppLayout>
        <div className="max-w-3xl mx-auto">
          <div className="bg-white rounded-lg shadow-md p-8">
            <div className="animate-pulse space-y-4">
              <div className="h-8 bg-gray-200 rounded w-1/3"></div>
              <div className="h-10 bg-gray-200 rounded"></div>
              <div className="h-64 bg-gray-200 rounded"></div>
            </div>
          </div>
        </div>
      </AppLayout>
    );
  }

  return (
    <AppLayout>
      <div className="max-w-3xl mx-auto">
        <button
          onClick={handleCancel}
          className="inline-flex items-center gap-2 text-sm text-gray-600 hover:text-gray-900 mb-6"
        >
          <ArrowLeft className="w-4 h-4" />
          Back to Notes
        </button>

        <div className="bg-white rounded-lg shadow-md p-8">
          <h1 className="text-2xl font-bold text-gray-900 mb-6">Edit Note</h1>

          <form onSubmit={handleSubmit} className="space-y-6">
            <Input
              type="text"
              name="title"
              label="Title"
              placeholder="Enter note title"
              value={formData.title}
              onChange={handleChange}
              error={errors.title}
            />

            <Textarea
              name="content"
              label="Content"
              placeholder="Write your note content here..."
              value={formData.content}
              onChange={handleChange}
              error={errors.content}
              rows={12}
            />

            <div className="flex gap-3 justify-end">
              <Button
                type="button"
                variant="secondary"
                onClick={handleCancel}
                disabled={isSaving}
              >
                Cancel
              </Button>
              <Button
                type="submit"
                variant="primary"
                isLoading={isSaving}
              >
                Update Note
              </Button>
            </div>
          </form>
        </div>
      </div>
    </AppLayout>
  );
};
