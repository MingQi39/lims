import client from './client'
import type { CreateNoteReq, Note, PageData, UpdateNoteReq } from './types'

export interface NoteListParams {
  page?: number
  page_size?: number
  keyword?: string
}

export const notesApi = {
  list: (params: NoteListParams) =>
    client.get<never, PageData<Note>>('/api/v1/notes', { params }),

  create: (data: CreateNoteReq) =>
    client.post<never, Note>('/api/v1/notes', data),

  update: (id: number, data: UpdateNoteReq) =>
    client.put<never, void>(`/api/v1/notes/${id}`, data),

  remove: (id: number, updatedAt: string) =>
    client.delete<never, void>(`/api/v1/notes/${id}`, {
      params: { updated_at: updatedAt },
    }),
}
