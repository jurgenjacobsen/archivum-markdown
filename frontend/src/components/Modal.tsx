import { X } from 'lucide-react';

interface ModalProps {
  isOpen: boolean;
  onClose: () => void;
  onConfirm?: () => void;
  title: string;
  message: string;
  confirmLabel?: string;
  cancelLabel?: string;
  type?: 'alert' | 'confirm';
}

export const Modal = ({ 
  isOpen, 
  onClose, 
  onConfirm, 
  title, 
  message, 
  confirmLabel = "Confirm", 
  cancelLabel = "Cancel",
  type = 'confirm'
}: ModalProps) => {
  if (!isOpen) return null;

  return (
    <div className="fixed inset-0 z-[100] flex items-center justify-center p-4 bg-[#242424]/75 backdrop-blur-sm no-print">
      <div className="bg-white border-[4px] border-[#242424] w-full max-w-md animate-modal shadow-[12px_12px_0_0_rgba(0,0,0,1)]">
        {/* Header */}
        <div className="bg-[#242424] text-white px-6 py-3 flex justify-between items-center">
          <h3 className="text-xs uppercase font-black tracking-[0.2em]">{title}</h3>
          <button onClick={onClose} className="hover:rotate-90 transition-transform duration-300">
            <X size={18} />
          </button>
        </div>

        {/* Content */}
        <div className="p-8">
          <p className="text-sm font-bold leading-relaxed text-[#242424] uppercase tracking-tight">
            {message}
          </p>
        </div>

        {/* Footer */}
        <div className="p-6 pt-0 flex justify-end space-x-4">
          <button 
            onClick={onClose}
            className="px-6 py-2 text-[10px] uppercase font-black border-2 border-[#242424] hover:bg-gray-100 transition-colors tracking-widest"
          >
            {cancelLabel}
          </button>
          
          {type === 'confirm' && onConfirm && (
            <button 
              onClick={() => {
                onConfirm();
                onClose();
              }}
              className="px-6 py-2 text-[10px] uppercase font-black bg-[#242424] text-white border-2 border-[#242424] hover:bg-[#242424]/75 transition-colors tracking-widest"
            >
              {confirmLabel}
            </button>
          )}
        </div>
      </div>
    </div>
  );
};
