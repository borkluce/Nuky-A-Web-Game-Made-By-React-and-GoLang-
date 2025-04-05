import { createClient } from '@supabase/supabase-js'

const supabaseUrl = "https://chhtdehudjctxfvsvtfl.supabase.co"
const supabaseAnonKey = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJzdXBhYmFzZSIsInJlZiI6ImNoaHRkZWh1ZGpjdHhmdnN2dGZsIiwicm9sZSI6ImFub24iLCJpYXQiOjE3NDM1MjA4NTMsImV4cCI6MjA1OTA5Njg1M30.QAuBv9s-9uyCDMAUqjCBY_GBWOqEoL1QgUChwcwPuGM"

export const supabase = createClient(supabaseUrl, supabaseAnonKey)
