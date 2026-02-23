class PaginationEntity<T> {
  final List<T> data;
  final int page;
  final int pageSize;
  final int totalItems;
  final int totalPages;
  final bool hasNextPage;
  final bool hasPrevPage;

  PaginationEntity({
    required this.data,
    required this.page,
    required this.pageSize,
    required this.totalItems,
    required this.totalPages,
    required this.hasNextPage,
    required this.hasPrevPage,
  });
}