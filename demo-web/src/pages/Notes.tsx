import { useMemo, useState } from 'react'
import { useAuthStore } from '@/stores/authStore'
import {
  useCreateNote,
  useDeleteNote,
  useNotesList,
  useUpdateNote,
} from '@/hooks/useNotes'
import type { Note } from '@/api/types'

export default function NotesPage() {
  const userName = useAuthStore((s) => s.userName)
  const logout = useAuthStore((s) => s.logout)
  const [page, setPage] = useState(1)
  const [keyword, setKeyword] = useState('')
  const [search, setSearch] = useState('')
  const [editing, setEditing] = useState<Note | null>(null)
  const [title, setTitle] = useState('')
  const [content, setContent] = useState('')

  const listQuery = useNotesList({ page, page_size: 10, keyword: search })
  const createMut = useCreateNote()
  const updateMut = useUpdateNote()
  const deleteMut = useDeleteNote()

  const totalPages = useMemo(() => {
    const total = listQuery.data?.total ?? 0
    return Math.max(1, Math.ceil(total / 10))
  }, [listQuery.data?.total])

  const resetForm = () => {
    setEditing(null)
    setTitle('')
    setContent('')
  }

  const startEdit = (note: Note) => {
    setEditing(note)
    setTitle(note.title)
    setContent(note.content)
  }

  const handleSave = () => {
    const trimmed = title.trim()
    if (!trimmed) {
      alert('请输入标题')
      return
    }
    if (editing) {
      updateMut.mutate(
        {
          id: editing.id,
          data: {
            title: trimmed,
            content,
            updated_at: editing.updated_at,
          },
        },
        { onSuccess: resetForm },
      )
      return
    }
    createMut.mutate(
      { title: trimmed, content },
      { onSuccess: resetForm },
    )
  }

  return (
    <div className="mx-auto max-w-3xl p-6">
      <header className="mb-6 flex flex-wrap items-center justify-between gap-3">
        <div>
          <h1 className="text-2xl font-semibold">我的笔记</h1>
          <p className="text-sm text-slate-500">你好，{userName ?? '用户'}</p>
        </div>
        <button
          type="button"
          onClick={logout}
          className="rounded-lg border border-slate-300 px-3 py-1.5 text-sm hover:bg-white"
        >
          退出
        </button>
      </header>

      <section className="mb-6 rounded-xl border border-slate-200 bg-white p-4 shadow-sm">
        <h2 className="text-sm font-medium text-slate-700">
          {editing ? '编辑笔记' : '新建笔记'}
        </h2>
        <input
          className="mt-3 w-full rounded-lg border border-slate-300 px-3 py-2 text-sm"
          placeholder="标题"
          value={title}
          onChange={(e) => setTitle(e.target.value)}
        />
        <textarea
          className="mt-2 w-full rounded-lg border border-slate-300 px-3 py-2 text-sm"
          placeholder="内容（可选）"
          rows={3}
          value={content}
          onChange={(e) => setContent(e.target.value)}
        />
        <div className="mt-3 flex gap-2">
          <button
            type="button"
            onClick={handleSave}
            disabled={createMut.isPending || updateMut.isPending}
            className="rounded-lg bg-slate-900 px-4 py-2 text-sm text-white hover:bg-slate-800 disabled:opacity-60"
          >
            {editing ? '保存' : '创建'}
          </button>
          {editing && (
            <button
              type="button"
              onClick={resetForm}
              className="rounded-lg border border-slate-300 px-4 py-2 text-sm"
            >
              取消
            </button>
          )}
        </div>
      </section>

      <section className="mb-4 flex gap-2">
        <input
          className="flex-1 rounded-lg border border-slate-300 px-3 py-2 text-sm"
          placeholder="搜索标题或内容"
          value={keyword}
          onChange={(e) => setKeyword(e.target.value)}
        />
        <button
          type="button"
          className="rounded-lg bg-slate-800 px-4 py-2 text-sm text-white"
          onClick={() => {
            setSearch(keyword.trim())
            setPage(1)
          }}
        >
          搜索
        </button>
      </section>

      {listQuery.isLoading && (
        <p className="text-sm text-slate-500">加载中…</p>
      )}
      {listQuery.isError && (
        <p className="text-sm text-red-600">{listQuery.error.message}</p>
      )}

      <ul className="space-y-3">
        {(listQuery.data?.list ?? []).map((note) => (
          <li
            key={note.id}
            className="rounded-xl border border-slate-200 bg-white p-4 shadow-sm"
          >
            <div className="flex items-start justify-between gap-2">
              <div>
                <h3 className="font-medium">{note.title}</h3>
                {note.content && (
                  <p className="mt-1 whitespace-pre-wrap text-sm text-slate-600">
                    {note.content}
                  </p>
                )}
                <p className="mt-2 text-xs text-slate-400">
                  更新于 {new Date(note.updated_at).toLocaleString()}
                </p>
              </div>
              <div className="flex shrink-0 gap-2">
                <button
                  type="button"
                  className="text-sm text-slate-700 hover:underline"
                  onClick={() => startEdit(note)}
                >
                  编辑
                </button>
                <button
                  type="button"
                  className="text-sm text-red-600 hover:underline"
                  onClick={() => {
                    if (!confirm('确定删除这条笔记？')) return
                    deleteMut.mutate({ id: note.id, updatedAt: note.updated_at })
                  }}
                >
                  删除
                </button>
              </div>
            </div>
          </li>
        ))}
      </ul>

      {!listQuery.isLoading && (listQuery.data?.list?.length ?? 0) === 0 && (
        <p className="text-sm text-slate-500">暂无笔记，先创建一条吧。</p>
      )}

      <div className="mt-6 flex items-center justify-between text-sm">
        <button
          type="button"
          disabled={page <= 1}
          className="rounded border px-3 py-1 disabled:opacity-40"
          onClick={() => setPage((p) => Math.max(1, p - 1))}
        >
          上一页
        </button>
        <span>
          第 {page} / {totalPages} 页（共 {listQuery.data?.total ?? 0} 条）
        </span>
        <button
          type="button"
          disabled={page >= totalPages}
          className="rounded border px-3 py-1 disabled:opacity-40"
          onClick={() => setPage((p) => p + 1)}
        >
          下一页
        </button>
      </div>
    </div>
  )
}
