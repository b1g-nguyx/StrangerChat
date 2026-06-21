import { AuthGuard } from '@/features/auth';

export default function Home() {
  return (
    <AuthGuard>
      <div className="flex flex-col flex-1 items-center justify-center text-center p-4">
        <h1 className="text-3xl font-bold text-slate-50 mb-4">
          Welcome to Stranger Chat
        </h1>
        <p className="text-slate-400 max-w-md">
          You are successfully logged in! The lobby and random matching features will be implemented here.
        </p>
      </div>
    </AuthGuard>
  );
}
