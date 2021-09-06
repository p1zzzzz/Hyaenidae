export default {
  data() {
    return {
      page: 1,
      total: 10,
      pageSize: 10,
      tableData: [],
      searchInfo: {}
    }
  },
  methods: {
    filterDict(value, type) {
      const rowLabel = this[type + 'Options'] && this[type + 'Options'].filter(item => item.value === value)
      return rowLabel && rowLabel[0] && rowLabel[0].label
    },
    handleSizeChange(val) {
      this.pageSize = val
      this.getTableData()
    },
    handleCurrentChange(val) {
      this.page = val
      this.getTableData()
    },
    async getTableData(page = this.page, pageSize = this.pageSize,searchkey=this.searchkey) {
      const table = await this.listApi({ page, pageSize, searchkey })
      if (table.code === 0) {
        this.tableData = table.data.list
        this.total = table.data.total
        this.page = table.data.page
        this.pageSize = table.data.pageSize
      }
    }
  }
}
