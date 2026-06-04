import { useMutation, useQuery, useQueryClient } from '@tanstack/react-query'
import { notesApi, type NoteListParams } from '@/api/notes'
import type { CreateNoteReq, UpdateNoteReq } from '@/api/types'
import { getApiErrorMessage } from '@/api/client'

const notesKey = ['notes'] as const

export function useNotesList(params: NoteListParams) {
  return useQuery({
    queryKey: [...notesKey, params],
    queryFn: () => notesApi.list(params),
  })
}

export function useCreateNote() {
  const qc = useQueryClient()
  return useMutation({
    mutationFn: (data: CreateNoteReq) => notesApi.create(data),
    onSuccess: () => qc.invalidateQueries({ queryKey: notesKey }),
    onError: (err) => alert(getApiErrorMessage(err)),
  })
}

export function useUpdateNote() {
  const qc = useQueryClient()
  return useMutation({
    mutationFn: ({ id, data }: { id: number; data: UpdateNoteReq }) =>
      notesApi.update(id, data),
    onSuccess: () => qc.invalidateQueries({ queryKey: notesKey }),
    onError: (err) => alert(getApiErrorMessage(err)),
  })
}

export function useDeleteNote() {
  const qc = useQueryClient()
  return useMutation({
    mutationFn: ({ id, updatedAt }: { id: number; updatedAt: string }) =>
      notesApi.remove(id, updatedAt),
    onSuccess: () => qc.invalidateQueries({ queryKey: notesKey }),
    onError: (err) => alert(getApiErrorMessage(err)),
  })
}
