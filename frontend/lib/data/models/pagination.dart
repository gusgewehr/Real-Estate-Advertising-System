
import '../../domain/entity/pagination.dart';

class PaginationModel<T> extends PaginationEntity<T> {

  PaginationModel({
    required super.data,
    required super.page,
    required super.pageSize,
    required super.totalItems,
    required super.totalPages,
    required super.hasNextPage,
    required super.hasPrevPage,

});

  factory PaginationModel.fromJson(Map<String, dynamic> json,
      T Function(Map<String, dynamic>) creator,) {

    return PaginationModel(
      data: (json['data'] as List)
          .map((item) => creator(item as Map<String, dynamic>))
          .toList(),
      page: json['page'] ?? 0,
      pageSize: json['page_size'] ?? 0,
      totalItems: json['total_items'] ?? 0,
      totalPages: json['total_pages'] ?? 0,
      hasNextPage: json['has_next_page'] ?? false,
      hasPrevPage: json['has_prev_page'] ?? false,
    );
  }




}