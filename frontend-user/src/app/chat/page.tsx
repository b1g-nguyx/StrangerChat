import { ChatView } from '@/features/chat/components/ChatView';
import { AuthGuard } from '@/features/auth/components/auth-guard';

export default function ChatPage() {
  return (
    <AuthGuard>
      <ChatView />
    </AuthGuard>
  );
}
