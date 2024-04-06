export interface ResponseError {
  code: number;
  message: string;
  data: {};
}

export interface BaseResponse {
  page: number;
  perPage: number;
  totalPages: number;
  totalItems?: number;
  items: any[];
}
