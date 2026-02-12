import React, { useState } from 'react';
import { useNavigate } from 'react-router';
import { ArrowLeft } from 'lucide-react';
import { AppLayout } from '../layouts/AppLayout';
import { Input } from '../components/Input';
import { Textarea } from '../components/Textarea';
import { Button } from '../components/Button';
import { notesService } from '../services/notesService';
import { toast } from 'sonner@2.0.3';

export const CreateNotePage: React.FC = () => {
  const navigate = useNavigate();
  const [formData, setFormData] = useState({
    title: '',
    content: '',
  });
  const [errors, setErrors] = useState<Record<string, string>>({});
  const [isLoading, setIsLoading] = useState(false);

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

    if (!validate()) return;

    setIsLoading(true);

    try {
      await notesService.createNote({
        title: formData.title,
        content: formData.content,
      });

      toast.success('Note created');
      navigate('/notes');
    } catch (error: any) {
      toast.error(error.message || 'Failed to create note');
    } finally {
      setIsLoading(false);
    }
  };

  const handleCancel = () => {
    navigate('/notes');
  };

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
          <h1 className="text-2xl font-bold text-gray-900 mb-6">Create New Note</h1>

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
                disabled={isLoading}
              >
                Cancel
              </Button>
              <Button
                type="submit"
                variant="primary"
                isLoading={isLoading}
              >
                Save Note
              </Button>
            </div>
          </form>
        </div>
      </div>
    </AppLayout>
  );
};
