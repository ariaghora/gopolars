#ifndef LIBCPOLARS_H_
#define LIBCPOLARS_H_

#include <stddef.h>
#include <stdlib.h>

typedef struct {
        void* data;
        int   result_code;
        char* message;
} DataFrameResult;

void**          lcp_alloc_expr_arr(int32_t len);
DataFrameResult lcp_collect(void* lf);
DataFrameResult lcp_columns(void* df, char** names, size_t len);
char*           lcp_dataframe_to_str(void* df);
void*           lcp_expr_alias(void* expr, const char* name);
void*           lcp_expr_column(const char* name);
void*           lcp_expr_eq(void* e1, void* e2);
void*           lcp_expr_i32(int32_t val);
void*           lcp_expr_str(const char* val);
void*           lcp_filter(void* lf, void* predicate);
int             lcp_frame_equal(void* df1, void* df2, int missing);
void*           lcp_head(void* df1, int n);
void*           lcp_lazy(void* df);
void*           lcp_new_int_series(char* name, int* data, size_t size);
DataFrameResult lcp_read_csv(char* path);
char*           lcp_series_to_str(void* s);
void*           lcp_select(void* lf, void* const* exprs, int32_t len);
void            lcp_set_expr_arr(void** expr_arr, void* expr, int32_t at);

#endif